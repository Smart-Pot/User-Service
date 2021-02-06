package data

import (
	"context"
	"fmt"
	"log"

	"userservice/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection of comments
var collection *mongo.Collection

// DatabaseConnection :
func DatabaseConnection() {
	fmt.Println("Connected to Database!")
	clientOptions := options.Client().ApplyURI(config.C.Database.Addr)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(config.C.Database.DBName).Collection("users")

}
