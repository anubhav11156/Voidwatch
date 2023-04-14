package main

import (
	"fmt"
	"net/http"
	"log"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)


const port = 8080
const atlasConnectionUri = "mongodb+srv://anubhav11697:myMongo123@myfirstcluster.hfdwigv.mongodb.net/?retryWrites=true&w=majority"

type application struct {
	Domain string
}

func main(){
	// steps
	// 1. set application configs
	// 2. read from command line
	// 3. connnect to database
	// 4. start a web server

	var app application

	app.Domain = "example.com"

	log.Println("Starting application at port : ",port)

	// http.HandleFunc("/",Hello)

	// starts a web server
	err := http.ListenAndServe(fmt.Sprintf("https://subwatch-backend.onrender.com:%d",port), app.routes())
	if err != nil {
		log.Fatal(err)
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

	// check the connection

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("DB connection succeeded!")
	}

	db := client.Database("myDatabase")

	return db, nil
}
