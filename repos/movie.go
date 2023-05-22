package repos

import (
	"database/sql"

	"github.com/dunzoit/BookMyShow/dtos"
	"github.com/dunzoit/BookMyShow/models"
)

const (
	getMovies    = `SELECT id, name, cost, description, poster, trailer, duration, rating, fun_facts, grade FROM movies`
	addMovie     = `INSERT INTO movies (name, cost, description, poster, trailer, duration, rating, fun_facts, grade) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	getMovieByID = `SELECT id, name, cost, description, poster, trailer, duration, rating, fun_facts, grade FROM movies WHERE id = $1`
	updateMovie  = `UPDATE movies SET name = $1, cost = $2, description = $3, poster = $4, trailer = $5, duration = $6, rating = $7, fun_facts = $8, grade = $9 WHERE id = $10`
	deleteMovie  = `DELETE FROM movies WHERE id = $1`
)

type Movie struct {
	db *sql.DB
}

func NewMovieRepo(db *sql.DB) *Movie {
	return &Movie{
		db: db,
	}
}

func (m *Movie) GetMovies() ([]*models.Movie, error) {
	rows, err := m.db.Query(getMovies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*models.Movie
	for rows.Next() {
		var id int
		var name string
		var cost float64
		var description string
		var poster string
		var trailer string
		var duration int
		var rating float64
		var funFacts string
		var grade string
		if err := rows.Scan(&id, &name, &cost, &description, &poster, &trailer, &duration, &rating, &funFacts,
			&grade); err != nil {
			return nil, err
		}
		movie := &models.Movie{
			ID:          id,
			Name:        name,
			Cost:        cost,
			Description: description,
			Poster:      poster,
			Trailer:     trailer,
			Duration:    duration,
			Rating:      rating,
			FunFacts:    funFacts,
			Grade:       grade,
		}
		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}

func (m *Movie) AddMovie(movie *dtos.CreateMovieRequest) error {
	_, err := m.db.Exec(addMovie, movie.Name, movie.Cost, movie.Description, movie.Poster, movie.Trailer,
		movie.Duration, movie.Rating, movie.FunFacts, movie.Grade)
	if err != nil {
		return err
	}
	return nil
}

func (m *Movie) GetMovieByID(movieID int) (*models.Movie, error) {
	row := m.db.QueryRow(getMovieByID, movieID)

	var id int
	var name string
	var cost float64
	var description string
	var poster string
	var trailer string
	var duration int
	var rating float64
	var funFacts string
	var grade string
	if err := row.Scan(&id, &name, &cost, &description, &poster, &trailer, &duration, &rating, &funFacts,
		&grade); err != nil {
		return nil, err
	}

	movie := &models.Movie{
		ID:          id,
		Name:        name,
		Cost:        cost,
		Description: description,
		Poster:      poster,
		Trailer:     trailer,
		Duration:    duration,
		Rating:      rating,
		FunFacts:    funFacts,
		Grade:       grade,
	}

	return movie, nil
}

func (m *Movie) UpdateMovie(movieID int, updatedValues *models.Movie) error {
	_, err := m.db.Exec(updateMovie, updatedValues.Name, updatedValues.Cost, updatedValues.Description,
		updatedValues.Poster, updatedValues.Trailer, updatedValues.Duration, updatedValues.Rating,
		updatedValues.FunFacts, updatedValues.Grade, movieID)
	if err != nil {
		return err
	}
	return nil
}

func (m *Movie) DeleteMovie(movieID int) error {
	_, err := m.db.Exec(deleteMovie, movieID)
	if err != nil {
		return err
	}
	return nil
}
