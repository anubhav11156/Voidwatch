package respository

import "backend/internal/models"

type DatabaseRepo interface {
	AllMovies() ([]*models.Movie, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(userId int) (*models.User, error)

	GetOneMovie(id int) (*models.Movie, error)
}
