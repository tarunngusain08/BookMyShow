package services

import (
	"github.com/dunzoit/BookMyShow/models"
	"github.com/dunzoit/BookMyShow/repos"
)

type Auditorium struct {
	repo repos.AuditoriumRepository
}

func NewAuditoriumService(repo repos.AuditoriumRepository) *Auditorium {
	return &Auditorium{
		repo: repo,
	}
}

func (c *Auditorium) GetAuditoriums() ([]*models.Auditorium, error) {
	return c.repo.GetAuditoriums()
}

func (c *Auditorium) AddAuditorium(Auditoriums *models.Auditorium) error {
	return c.repo.AddAuditorium(Auditoriums)
}

func (c *Auditorium) GetAuditorium(AuditoriumId int) (*models.Auditorium, error) {
	return c.repo.GetAuditorium(AuditoriumId)
}

func (c *Auditorium) UpdateAuditorium(AuditoriumId int, updatedValues *models.Auditorium) error {
	return c.repo.UpdateAuditorium(AuditoriumId, updatedValues)
}

func (c *Auditorium) DeleteAuditorium(AuditoriumId int) error {
	return c.repo.DeleteAuditorium(AuditoriumId)
}
