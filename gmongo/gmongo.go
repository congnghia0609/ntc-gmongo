package gmongo

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MClient *mongo.Client
var TestDB *mongo.Database

func InitMongo() {
	// Set client options
	// https://docs.mongodb.com/manual/reference/connection-string/
	clientOptions := options.Client().ApplyURI("mongodb://nghiatc:pwtest123@localhost:27017/testdb?authSource=admin&retryWrites=true&retryReads=true&maxPoolSize=100&maxIdleTimeMS=60000")

	// Connect to MongoDB
	var err error
	MClient, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = MClient.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	// new TestDB
	TestDB = MClient.Database("testdb")

	fmt.Println("Connected to MongoDB!")
}

func MClose() {
	MClient.Disconnect(context.TODO())
	fmt.Println("Disconnect MongoDB!")
}
