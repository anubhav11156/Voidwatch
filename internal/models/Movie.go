package models

// "go.mongodb.org/mongo-driver/bson"

type Movie struct {
	ID          string `json:"id" bson:"_id"`
	Title       string `json:"title" bson:"title"`
	ReleaseDate string `json:"release_date" bson:"release_date"`
	RunTime     string `json:"runtime" bson:"runtime"`
	MPAARating  string `json:"mpaa_rating" bson:"mpaa_rating"`
	Description string `json:"description" bson:"description"`
	Image       string `json:"image" bson:"image"`
	CreatedAt   string `json:"createdAt" bson:"createdAt"`
	UpdatedAt   string `json:"updatedAt" bson:"updatedAt"`
}
