package Database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// set connection to database
func DatabaseConnection() *mongo.Client {
	//creating a context to timeout after 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	
	//connecting to the database here
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		//falling to connect to the database will lead us here
		log.Fatal(err)
	}

	//this will be called when the connection timeouts
	defer cancel()

	//verifying the connection
	err = client.Ping(context.TODO(), nil)
	
	//falied to connect to the database
	if err != nil {
		log.Println("Unable to connect to database")
		return nil
	}

	//connected to the database and returning the client
	fmt.Println("Connected to database" )
	return client
}

var Client *mongo.Client = DatabaseConnection()

// get user data
func UserData(client mongo.Client, collectionName string) *mongo.Collection {

}

// get product data
func ProductData(client mongo.Client, collectionName string) *mongo.Collection {
	
}