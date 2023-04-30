package models

// "go.mongodb.org/mongo-driver/bson"

type Movie struct {
	ID          int      `json:"_id" bson:"_id"`
	Title       string   `json:"title" bson:"title"`
	ReleaseDate string   `json:"release_date" bson:"release_date"`
	RunTime     string   `json:"runtime" bson:"runtime"`
	MPAARating  string   `json:"mpaa_rating" bson:"mpaa_rating"`
	Description string   `json:"description" bson:"description"`
	Image       string   `json:"image" bson:"image"`
	CreatedAt   string   `json:"createdAt" bson:"createdAt"`
	UpdatedAt   string   `json:"updatedAt" bson:"updatedAt"`
	Genres      []*Genre `json:"genres" bson:"genres"`
	GenresArray []int    `json:"genres_array" bson:"genres_array"`
}

type Genre struct {
	ID        int    `json:"_id" bson:"_id"`
	Genre     string `json:"genre" bson:"genre"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
	Checked   bool   `json:"checked" bson:"checked"`
}

type MovieGenres struct {
	ID      int `json:"_id" bson:"_id"`
	MovieId int `json:"movie_id" bson:"movie_id"`
	GenreId int `json:"genre_id" bson:"genre_id"`
}
