package services

import (
	"github.com/dunzoit/BookMyShow/dtos"
	"github.com/dunzoit/BookMyShow/models"
	"github.com/dunzoit/BookMyShow/repos"
)

type Movie struct {
	repo repos.MovieRepository
}

func NewMovieService(repo repos.MovieRepository) *Movie {
	return &Movie{
		repo: repo,
	}
}

func (c *Movie) GetMovies() ([]*models.Movie, error) {
	return c.repo.GetMovies()
}

func (c *Movie) AddMovie(Movies *dtos.CreateMovieRequest) error {
	return c.repo.AddMovie(Movies)
}

func (c *Movie) GetMovie(MovieId int) (*models.Movie, error) {
	return c.repo.GetMovieByID(MovieId)
}

func (c *Movie) UpdateMovie(MovieId int, updatedValues *models.Movie) error {
	return c.repo.UpdateMovie(MovieId, updatedValues)
}

func (c *Movie) DeleteMovie(MovieId int) error {
	return c.repo.DeleteMovie(MovieId)
}
