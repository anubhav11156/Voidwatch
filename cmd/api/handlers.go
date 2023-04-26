package main

import (
	"context"
	"errors"
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
	// read json payload submitted by the end user
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate user against database
	user, err := app.Database.GetUserByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("Invalid credentials"), http.StatusBadRequest)
		return
	}

	// check password (match the hash)
	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// create a jwt user
	u := jwtUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	// generate tokens
	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	refreshCookie := app.auth.GetRefreshCookie(tokens.Token)
	http.SetCookie(w, refreshCookie)

	app.writeJSON(w, http.StatusAccepted, tokens)

}
