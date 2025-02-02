package repository

import (
	"backend/internal/models"
	"database/sql"
)

type DatabaseRepo interface {
	//connection
	Connection() *sql.DB
	// get all movie
	AllMovies() ([]*models.Movie, error)
	OneMovie(id int) (*models.Movie, error)
	OneMovieForEdit(id int) (*models.Movie, []*models.Genre, error)

	//get user by email
	GetUserByEmail(email string) (*models.User, error)
	//get use by id
	GetUserByID(id int) (*models.User, error)
}
