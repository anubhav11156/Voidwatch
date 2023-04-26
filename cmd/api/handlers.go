package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	var payload = struct {
		Status      string   `json:"status"`
		Message     string   `json:"message"`
		Version     string   `json:"version"`
		Collections []string `josn:"collections"`
	}{
		Status:  "active",
		Message: "Welcome to subwatch!",
		Version: "1.0.0",
	}

	db, err := app.connectToDB()
	if err != nil {
		log.Println(err)
	}

	// get all collections in you database
	collections, err := db.ListCollectionNames(context.TODO(), bson.M{})

	if err != nil {
		fmt.Println(err)
	} else {

		for _, coll := range collections {
			payload.Collections = append(payload.Collections, coll)
		}
		_ = app.writeJSON(w, http.StatusOK, payload)
	}

}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {

	// app.Database is alredy initalized to mongodb connection
	movies, err := app.Database.AllMovies()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	_ = app.writeJSON(w, http.StatusOK, movies)
}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	// read json payload

	// validate user against database

	// chekc password (mathc the hash)

	// create a jwt user
	u := jwtUser{
		ID:        1,
		FirstName: "Anubhav",
		LastName:  "Kumar",
	}

	// generate tokens
	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	w.Write([]byte(tokens.Token))

}
