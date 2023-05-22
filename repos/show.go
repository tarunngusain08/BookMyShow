package repos

import (
	"database/sql"
	"time"

	"github.com/dunzoit/BookMyShow/dtos"
	"github.com/dunzoit/BookMyShow/models"
)

const (
	getShows    = `SELECT id, start_time, end_time, theatre_id, screen, movie_id, features FROM shows`
	addShow     = `INSERT INTO shows (start_time, end_time, theatre_id, screen, movie_id, features) VALUES ($1, $2, $3, $4, $5, $6)`
	getShowByID = `SELECT id, start_time, end_time, theatre_id, screen, movie_id, features FROM shows WHERE id = $1`
	updateShow  = `UPDATE shows SET start_time = $1, end_time = $2, theatre_id = $3, screen = $4, movie_id = $5, features = $6 WHERE id = $7`
	deleteShow  = `DELETE FROM shows WHERE id = $1`
)

type Show struct {
	db *sql.DB
}

func NewShowRepo(db *sql.DB) *Show {
	return &Show{
		db: db,
	}
}

func (s *Show) GetShows() ([]*models.Show, error) {
	rows, err := s.db.Query(getShows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shows []*models.Show
	for rows.Next() {
		var id int
		var startTime, endTime time.Time
		var theatreID, screen, movieID int
		var features []string
		if err := rows.Scan(&id, &startTime, &endTime, &theatreID, &screen, &movieID, &features); err != nil {
			return nil, err
		}
		show := &models.Show{
			ID:        id,
			StartTime: &startTime,
			EndTime:   &endTime,
			Theatre:   &models.Theatre{ID: theatreID},
			Screen:    screen,
			Movie:     &models.Movie{ID: movieID},
			Features:  features,
		}
		shows = append(shows, show)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return shows, nil
}

func (s *Show) AddShow(show *dtos.CreateShowRequest) error {
	_, err := s.db.Exec(addShow, show.StartTime, show.EndTime, show.TheatreID, show.Screen, show.MovieID, show.Features)
	if err != nil {
		return err
	}
	return nil
}

func (s *Show) GetShowByID(showID int) (*models.Show, error) {
	row := s.db.QueryRow(getShowByID, showID)

	var id int
	var startTime, endTime time.Time
	var theatreID, screen, movieID int
	var features []string
	if err := row.Scan(&id, &startTime, &endTime, &theatreID, &screen, &movieID, &features); err != nil {
		return nil, err
	}

	show := &models.Show{
		ID:        id,
		StartTime: &startTime,
		EndTime:   &endTime,
		Theatre:   &models.Theatre{ID: theatreID},
		Screen:    screen,
		Movie:     &models.Movie{ID: movieID},
		Features:  features,
	}

	return show, nil
}

func (s *Show) UpdateShow(showID int, updatedValues *models.Show) error {
	_, err := s.db.Exec(updateShow, updatedValues.StartTime, updatedValues.EndTime, updatedValues.Theatre.ID,
		updatedValues.Screen, updatedValues.Movie.ID, updatedValues.Features, showID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Show) DeleteShow(showID int) error {
	_, err := s.db.Exec(deleteShow, showID)
	if err != nil {
		return err
	}
	return nil
}
