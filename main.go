package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/arun-kushwaha04/ecommerce-backend/Controllers"
	"github.com/arun-kushwaha04/ecommerce-backend/Database"
	"github.com/arun-kushwaha04/ecommerce-backend/Middlewares"
	"github.com/arun-kushwaha04/ecommerce-backend/Routes"
)

func main(){
	//starting the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := Controllers.NewApplication(Database.ProductData(Database.Client, "Products"), Database.UserData(Database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	Routes.UserRoute(router)
	router.Use(Middlewares.Authentication())

	router.GET("/addToCart", app.addToCart())
	router.GET("/removeItem", app.removeItem())
	router.GET("/cartCheckout", app.cartCheckout())
	router.GET("/instantBuy", app.instantBuy())

	log.Fatal(router.Run(":" + port))
}