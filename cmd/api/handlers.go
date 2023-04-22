package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	// "time"
	// "backend/internal/models"
	// "backend/internal/databases"
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

	// if db == nil {

	// 	var err error
	// 	dbErr := db.Client().Ping(context.TODO(), nil)

	// 	if err != nil {
	// 		log.Fatal(dbErr)
	// 	}

	// }

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

		out, err := json.Marshal(payload)

		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(out)
	}

}

// func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {

// 	var movies []models.Movie

// 	rd, _ := time.Parse("2006-01-02", "1986-03-07")

// 	Highlander := models.Movie{
// 		ID:          1,
// 		Title:       "Highlander",
// 		ReleaseDate: rd,
// 		MPAARating:  "R",
// 		RunTime:     116,
// 		Description: "A very nice movie",
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 	}

// 	movies = append(movies, Highlander)

// 	rd, _ = time.Parse("2006-01-02", "1981-06-12")

// 	rotla := models.Movie{
// 		ID:          2,
// 		Title:       "Raiders of the lost Art",
// 		ReleaseDate: rd,
// 		MPAARating:  "PG-13",
// 		RunTime:     115,
// 		Description: "Another very nice movie",
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 	}

// 	movies = append(movies, rotla)

// 	out, err := json.Marshal(movies)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(out)
// }

// func (app *application) createMovie(w http.ResponseWriter, r *http.Request) {

// 	// var movies models.Movie

// 	rd, _ := time.Parse("2006-01-02","1986-03-07")

// 	Highlander := models.Movie {
// 		ID:1,
// 		Title:"Highlander",
// 		ReleaseDate: rd,
// 		MPAARating: "R",
// 		RunTime: 116,
// 		Description: "A very nice movie",
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}

// 	client := databases.getMongoClient();

// 	collection := client.Database("myDatabase").Collection("myCollection")

// 	result, err := collection.InsertOne(context.TODO(), Highlander)

// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Inserted!")
// 	}

// }
