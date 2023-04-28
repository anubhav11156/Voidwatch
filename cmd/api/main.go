package main

import (
	"backend/internal/respository"
	"backend/internal/respository/dbrepo"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
	// "net"
)

const port = 8080

// var atlasConnectionUri string

type application struct {
	Domain       string
	Database     respository.DatabaseRepo
	auth         Auth
	JWTSecret    string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
}

func main() {
	// steps
	// 1. set application configs
	// 2. read from command line
	// 3. connnect to database
	// 4. start a web server
	var app application

	// read from command line
	flag.StringVar(&app.JWTSecret, "jwt-secret", "myJWTSecret", "sigining secret")
	flag.StringVar(&app.JWTIssuer, "jwt-issuer", "subwatch.com", "sigining issuer")
	flag.StringVar(&app.JWTAudience, "jwt-audience", "subwatch.com", "sigining audience")
	flag.StringVar(&app.CookieDomain, "cookie-domain", "localhost", "cookie domain")
	flag.StringVar(&app.Domain, "app-domain", "example.com", "domain")
	flag.Parse()
	// => connect to mongodb database
	connection, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	// app.Database is of type DatabaseRepo defined in repository, so u have to populate it wiht your connection
	app.Database = &dbrepo.MongoDBRepo{DB: connection}
	defer connection.Client().Disconnect(context.Background())

	app.auth = Auth{
		Issuer:        app.JWTIssuer,
		Audience:      app.JWTAudience,
		Secret:        app.JWTSecret,
		TokenExpiry:   time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookiePath:    "/",
		// under score unde score make it more secure
		CookieName:   "__Host-refresh-token",
		CookieDomain: app.CookieDomain,
	}

	log.Println("Starting application... ")
	// starts a web server
	serverErr := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), app.routes())
	// err := http.ListenAndServe(fmt.Sprintf("subwatch-backend.onrender.com:",port), app.routes())
	// err := http.ListenAndServe(net.JoinHostPort("0.0.0.0", port), app.routes())
	if err != nil {
		log.Fatal(serverErr)
	}
}
