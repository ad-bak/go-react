package repository

import (
	"backend/internal/models"
	"database/sql"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies(genre ...int) ([]*models.Movie, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
	OneMovie(id int) (*models.Movie, error)
	OneMovieForEdit(id int) (*models.Movie, []*models.Genre, error)
	AllGenres() ([]*models.Genre, error)
	InsertMovie(m models.Movie) (int, error)
	UpdateMovieGenres(id int, genreIds []int) error
	UpdateMovie(m models.Movie) error
	DeleteMovie(id int) error
}
