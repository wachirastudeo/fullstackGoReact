package dbrepo

import (
	"backend/internal/models"
	"context"
	"database/sql"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 5

//all movies return

// connection db
func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) AllMovies() ([]*models.Movie, error) {
	//context handle query time out
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)

	defer cancel()
	//query database
	query := `select id, title, release_date, runtime, mpaa_rating, description,coalesce(image,''), created_at, updated_at from movies order by title`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//declare movies slice

	var movies []*models.Movie
	//iterate through rows
	for rows.Next() {
		movie := &models.Movie{}
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.ReleaseDate,
			&movie.Runtime,
			&movie.MPAARating,
			&movie.Description,
			&movie.Image,
			&movie.Created_at,
			&movie.Updated_at,
		)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

// สร้างฟังชั่นค้นหา user จาก email

func (m *PostgresDBRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `select id,email,password,firstname,lastname,createdat,updatedat from users where email = $1`
	// prepare statement sql injection

	var user models.User
	row := m.DB.QueryRowContext(ctx, query, email)
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil

}
