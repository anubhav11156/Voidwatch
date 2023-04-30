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

// this for public user
func (m *MongoDBRepo) GetOneMovie(id int) (*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	collection := m.DB.Collection("movies")

	//apply a filter
	filter := bson.M{"_id": id}
	var movie models.Movie
	err := collection.FindOne(ctx, filter).Decode(&movie)
	if err != nil {
		return nil, err
	}

	// now get genre associated with this movie
	coll := m.DB.Collection("movies_genres")
	// apply a filter to get all the genres the movie is associate with
	fil := bson.M{"movie_id": id}
	cursor, err := coll.Find(ctx, fil)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var genres []*models.Genre
	for cursor.Next(context.Background()) {
		var currentGenre models.Genre
		err := cursor.Decode(&currentGenre)
		if err != nil {
			log.Fatal(err)
		}
		genres = append(genres, &currentGenre)
	}
	movie.Genres = genres

	return &movie, nil
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

func (m *MongoDBRepo) GetUserById(userId int) (*models.User, error) {
	// defined my own context
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	collection := m.DB.Collection("users")
	// apply a filter

	filter := bson.M{"_id": userId}
	var user models.User
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
