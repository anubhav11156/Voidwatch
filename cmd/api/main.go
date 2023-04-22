package main

import (
	"backend/internal/respository"
	"backend/internal/respository/dbrepo"
	"context"
	"fmt"
	"log"
	"net/http"
	// "net"
)

const port = 8080

// var atlasConnectionUri string

type application struct {
	Domain   string
	Database respository.DatabaseRepo
}

func main() {
	// steps
	// 1. set application configs
	// 2. read from command line
	// 3. connnect to database
	// 4. start a web server
	var app application

	app.Domain = "example.com"

	connection, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.Database = &dbrepo.MongoDBRepo{DB: connection}
	defer connection.Client().Disconnect(context.Background())

	log.Println("Starting application... ")

	// starts a web server
	serverErr := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), app.routes())
	// err := http.ListenAndServe(fmt.Sprintf("subwatch-backend.onrender.com:",port), app.routes())
	// err := http.ListenAndServe(net.JoinHostPort("0.0.0.0", port), app.routes())
	if err != nil {
		log.Fatal(serverErr)
	}
}
