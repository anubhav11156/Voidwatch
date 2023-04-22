package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "net"
)

const port = 8080

var atlasConnectionUri string

type application struct {
	Domain string
}

func main() {
	// steps
	// 1. set application configs
	// 2. read from command line
	// 3. connnect to database
	// 4. start a web server

	// load the env

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	atlasConnectionUri = os.Getenv("MONGO_DB")

	var app application

	app.Domain = "example.com"

	log.Println("Starting application... ")

	// starts a web server
	serverErr := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), app.routes())
	// err := http.ListenAndServe(fmt.Sprintf("subwatch-backend.onrender.com:",port), app.routes())
	// err := http.ListenAndServe(net.JoinHostPort("0.0.0.0", port), app.routes())
	if err != nil {
		log.Fatal(serverErr)
	}
}

// get your mongodb database
func GetDb() (*mongo.Database, error) {

	// clientOptions := options.ApplyURI(atlasConnectionUri)

	opts := options.Client().ApplyURI(atlasConnectionUri)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		fmt.Println(err)
		// return
	}

	// check connection of your database
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("DB connection succeeded!")
	}

	db := client.Database("myDatabase")

	return db, nil
}
