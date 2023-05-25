package services

import (
	"BookMyShow/models"
	"BookMyShow/repos"
)

type Seat struct {
	repo repos.SeatRepository
}

func NewSeatService(repo repos.SeatRepository) *Seat {
	return &Seat{
		repo: repo,
	}
}

func (c *Seat) GetSeats() ([]*models.Seat, error) {
	return c.repo.GetSeats()
}

func (c *Seat) AddSeat(Seats *models.Seat) error {
	return c.repo.AddSeat(Seats)
}

func (c *Seat) GetSeat(SeatId int) (*models.Seat, error) {
	return c.repo.GetSeatByID(SeatId)
}

func (c *Seat) UpdateSeat(SeatId int, updatedValues *models.Seat) error {
	return c.repo.UpdateSeat(SeatId, updatedValues)
}

func (c *Seat) DeleteSeat(SeatId int) error {
	return c.repo.DeleteSeat(SeatId)
}
