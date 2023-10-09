package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ConnectionString = "mongodb+srv://user:user@cluster0.tzw5ea6.mongodb.net/?retryWrites=true&w=majority"
const DBName = "netflix"
const CollectionName = "watchlist"

// most imp as a connection
var collection *mongo.Collection

// connect with mongo db
func init() {
	//client options
	clientOption := options.Client().ApplyURI(ConnectionString)

	// connect to mongo db
	client, err := mongo.Connect(context.TODO(), clientOption)
	errHandle(err)

	fmt.Println("MongoDB connection success")

	collection = client.Database(DBName).Collection(CollectionName)

	fmt.Println("Collection reference is readt ")

}

func errHandle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
