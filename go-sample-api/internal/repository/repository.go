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

	GetUserByEmail(email string) (*models.User, error)
}
