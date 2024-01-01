package Models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	
	ID                primitive.ObjectID     `json:"id" bson:"_id"`
	First_Name        *string                `json:"first_name" bson:"first_name"          validate:"required,min=2,max=30"`
	Last_Name         *string                `json:"last_name" bson:"last_name"            validate:"required,min=2,max=30"`
	Password          *string                `json:"password" bson:"password"              validate:"required,min=8"`
	Email             *string                `json:"email" bson:"email"                    validate:"required,email"` 
	Phone             *string	               `json:"phone" bson:"phone"                    validate:"required"`
	Token             *string   						 `json:"token" bson:"token"`
	Refresh_Token	    *string                `json:"refresh_token" bson:"refresh_token"`
	Created_At        time.Time              `json:"created_at" bson:"created_at"`
	Updated_At        time.Time              `json:"updated_at" bson:"updated_at"`
	User_ID           *string                `json:"user_id" bson:"user_id"`
	User_Cart         []ProductUser          `json:"user_cart" bson:"user_cart"`
	Address_Detail    []Address  	           `json:"address_detail" bson:"address_detail"`
	Order_Status      []Order                `json:"order_status" bson:"order_status"`
}

type Product struct {
	Product_ID        primitive.ObjectID     `json:"_id" bson:"_id"`
	Product_Name      *string                `json:"product_name" bson:"product_name"`
	Price             *float64               `json:"price" bson:"price"`
	Rating            *float64               `json:"rating" bson:"rating"`
	Image             *string                `json:"image" bson:"image"`
}

type ProductUser struct {
	Product_ID        primitive.ObjectID     `json:"_id" bson:"_id"`
	Product_Name      *string                `json:"product_name" bson:"product_name"`
	Price             *float64               `json:"price" bson:"price"`
	Rating            *float64               `json:"rating" bson:"rating"`
	Image             *string                `json:"image" bson:"image"`
}

type Address struct {
	Address_ID        primitive.ObjectID     `json:"_id" bson:"_id"`
	House             *string                `json:"house" bson:"house"`
	Street            *string                `json:"street" bson:"street"`
	City              *string                `json:"city" bson:"city"`
	Pincode           *string                `json:"pincode" bson:"pincode"`
}

type Order struct {
	Order_ID         primitive.ObjectID      `json:"_id" bson:"_id"`
	Order_Cart       []ProductUser           `json:"order_cart" bson:"order_cart"`
	Order_At         time.Time               `json:"order_at" bson:"order_at"`
	Price            *float64                `json:"price" bson:"price"`
	Discount         *float64                `json:"discount" bson:"discount"`
	Payment_Method   Payment                 `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital          bool                    `json:"digital" bson:"digital"`
	COD              bool	                   `json:"coding" bson:"coding"`
}