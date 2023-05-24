package services

import (
	"github.com/dunzoit/BookMyShow/dtos"
	"github.com/dunzoit/BookMyShow/models"
	"github.com/dunzoit/BookMyShow/repos"
)

type Show struct {
	repo repos.ShowRepository
}

func NewShowService(repo repos.ShowRepository) *Show {
	return &Show{
		repo: repo,
	}
}

func (c *Show) GetShows() ([]*models.Show, error) {
	return c.repo.GetShows()
}

func (c *Show) AddShow(Show *dtos.CreateShowRequest) error {
	return c.repo.AddShow(Show)
}

func (c *Show) GetShow(ShowId int) (*models.Show, error) {
	return c.repo.GetShowByID(ShowId)
}

func (c *Show) UpdateShow(ShowId int, updatedValues *models.Show) error {
	return c.repo.UpdateShow(ShowId, updatedValues)
}

func (c *Show) DeleteShow(ShowId int) error {
	return c.repo.DeleteShow(ShowId)
}
