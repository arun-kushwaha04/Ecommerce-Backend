package Controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/arun-kushwaha04/ecommerce-backend/Database"
	"github.com/arun-kushwaha04/ecommerce-backend/Models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection = Database.UserData(*Database.Client, "Users")
var ProductCollection *mongo.Collection = Database.ProductData(*Database.Client, "Products")
var Validate = validator.New()

// hashpassord
func HashPassowrd(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 32)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

// verifypassord
func VerifyPassord(userPassword string, givenPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(givenPassword))
	if err != nil{
		return false, "wrong password"
	}
	return true, "correct password"
}

// product view admin
func ProductViewAdmin() gin.HandlerFunc{

}

// search product
func SearchProduct() gin.HandlerFunc{
	return func (c *gin.Context) {
		// creating object to hold products
		var productList []Models.Product
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// querying the db to find all products
		cursor, err :=  ProductCollection.Find(ctx, bson.D{})
		if err != nil {
			log.Printf("err: searchProduct - error occured while searching for product, %s", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		// decoding cursor data in product list
		err = cursor.All(ctx, &productList)
		if err != nil {
			log.Printf("err: searchProduct - error occured while decoding cursor data in productlist object, %s", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		defer cursor.Close(ctx)
		if err := cursor.Err(); err != nil {
			log.Printf("err: searchProduct - error occured with cursor, %s", err)
			c.JSON(400, "Some error occured while parsing the data")
			return
		}

		c.JSON(200, gin.H{"data": productList})
	}
}

// search product by query
func SearchProductByQuery(query string) gin.HandlerFunc{
	return func (c *gin.Context) {
		var searchProduct []Models.Product
		var productSearchQuery = c.Query("name")

		if productSearchQuery == "" {
			log.Printf("info: searchProductByQuery - query string is empty")
			c.JSON(http.StatusBadRequest, "message: query string is empty")
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		cursor, err := ProductCollection.Find(ctx,bson.M{"product_name": bson.M{"$regex": productSearchQuery}})

		if err != nil {
			log.Printf("err: searchProductByQuery - error while fetching product, %s", err)
			c.JSON(http.StatusInternalServerError, gin.H{"err": "Internal Server Error"})
			return
		}

		err = cursor.All(ctx, &searchProduct)
		if err != nil {
			log.Printf("err: searchProductByQuery - error occured while decoding cursor data in productlist object, %s", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		defer cursor.Close(ctx)

		if err = cursor.Err(); err != nil {
			log.Printf("err: searchProductByQuery - error occured with cursor, %s", err)
			c.JSON(400, "Some error occured while parsing the data")
			return
		}

		c.JSON(200, gin.H{"data": searchProduct})
	}
}

// login
func Login() gin.HandlerFunc {
	return func (c *gin.Context){
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user Models.User
		var foundUser Models.User

		// deconding passed value in user object
		if err := c.BindJSON(&user); err != nil {
			log.Printf("err: login - unable to bind json to user object")
			c.JSON(http.StatusBadGateway, gin.H{"error": "Internal error occured"})
			return
		}

		// searchin for user in db using email id
		err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)

		if err != nil {
			log.Printf("err: login - error while searching for user in db")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		// verfying password provided by user
		isValidPassword, msg := VerifyPassord(*user.Password, *foundUser.Password)

		if !isValidPassword {
			log.Printf("err: login - %s", msg)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
			return
		}
		
		// TODO: add logic for generating user token
		// token,  refreshToken, _ := generate.TokenGenrator
		// generate.UpdateAllTokens
		c.JSON(http.StatusFound, gin.H{"data": foundUser})
	}
}

// signup
func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user Models.User

		if err := c.BindJSON(&user); err != nil {
			log.Printf("err: signup - unable to bind json to user object")
			c.JSON(http.StatusBadGateway, gin.H{"error": "Internal error occured"})
			return
		}

		// validating with the valdition defined in the model
		validationErr := Validate.Struct(user)

		if validationErr != nil {
			log.Printf("err: signup - unable to validate user struct")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect values for the user struct"})
			return
		}


		// checking if the user email and phone has already been stored in database
		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		
		if err != nil {
			log.Printf("err: signup - error occured while searching for user in database")
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		}

		count, err = UserCollection.CountDocuments(ctx, bson.M{"phone":  user.Phone})

		if err != nil {
			log.Printf("err: signup - error occured during searching for user in database")
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, bson.M{"error": "Bad Request"})
		}


		// hashing the password
		password := HashPassowrd(*user.Password)
		user.Password = &password

		// filling values to other user fields
		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()
		// TODO: add logic for generating user token
		// token,  refreshToken, _ := generate.TokenGenrator
		
		user.User_Cart = make([]Models.ProductUser, 0)
		user.Address_Detail = make([]Models.Address, 0)
		user.Order_Status = make([]Models.Order, 0)
		
		_, inserterr := UserCollection.InsertOne(ctx, user)

		if inserterr != nil { 
			log.Printf("err: signup - error while creating new user in database")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to register"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Successfully registered"})
	}
}