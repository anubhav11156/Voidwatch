package models

import (
	"time"
	// "go.mongodb.org/mongo-driver/bson"
)

type Movie struct {
	ID          int       `json:"id" bson:"_id"`
	Title       string    `json:"title" bson:"title"`
	ReleaseDate time.Time `json:"release_date" bson:"release_date"`
	RunTime     int       `json:"runtime" bson:"runtime"`
	MPAARating  string    `json:"mpaa_rating" bson:"mpaa_rating"`
	Description string    `json:"description" bson:"description"`
	Image       string    `json:"image" bson:"image"`
	CreatedAt   time.Time `json:"-" bson:"-"`
	UpdatedAt   time.Time `json:"-" bson:"-"`
}
