package dbrepo

import (
	"backend/internal/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBRepo struct {
	DB *mongo.Database
}

const dbTimeout = time.Second * 5 // 5 seconds to query the database

func (m *MongoDBRepo) AllMovies() ([]*models.Movie, error) {

	// defined my own context
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// here write the query
	collection := m.DB.Collection("movies")

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	var movies []*models.Movie

	for cursor.Next(context.Background()) {

		var currentMovie models.Movie
		err := cursor.Decode(&currentMovie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, &currentMovie)
	}

	return movies, nil
}

func (m *MongoDBRepo) GetUserByEmail(email string) (*models.User, error) {
	// defined my own context
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	collection := m.DB.Collection("users")
	// apply a filter

	filter := bson.M{"email": email}
	var user models.User
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
