package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var atlasConnectionUri string
var mongoClient *mongo.Client

func openDatabase() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	atlasConnectionUri = os.Getenv("MONGO_DB")

	// create a connection with connecton pooling
	clientOptions := options.Client().ApplyURI(atlasConnectionUri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping to verify connection is established or not
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}
	mongoClient = client
	log.Println("Connected to MongoDB!")
	return nil
}

func (app *application) connectToDB() (*mongo.Database, error) {

	if mongoClient == nil {
		err := openDatabase()
		if err != nil {
			return nil, err
		}
	}
	db := mongoClient.Database("myDatabase")

	return db, nil
}
