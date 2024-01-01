package Controllers

import "github.com/gin-gonic/gin"

// hashpassord

func HashPassowrd(password string) string {
	return ""
}

// verifypassord
func VerifyPassord(userPassword string, givenPassword string) (bool, string) {
	return false, ""
}

// product view admin
func ProductViewAdmin() gin.HandlerFunc{

}

// search product
func SearchProduct() gin.HandlerFunc{

}

// search product by query
func SearchProductByQuery(query string) gin.HandlerFunc{
	
}

// login
func Login() gin.HandlerFunc {

}

// signup
func Signup() gin.HandlerFunc {

}