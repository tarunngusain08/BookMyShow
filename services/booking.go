package services

import (
	"github.com/dunzoit/BookMyShow/models"
	"github.com/dunzoit/BookMyShow/repos"
)

type Booking struct {
	repo repos.BookingRepository
}

func NewBookingService(repo repos.BookingRepository) *Booking {
	return &Booking{
		repo: repo,
	}
}

func (c *Booking) GetBookings(userId int) ([]*models.Booking, error) {
	return c.repo.GetBookings(userId)
}

func (c *Booking) MakeBooking(Booking *models.Booking) error {
	return c.repo.MakeBooking(Booking)
}

func (c *Booking) GetBooking(BookingId int) (*models.Booking, error) {
	return c.repo.GetBookingByID(BookingId)
}

func (c *Booking) UpdateBooking(BookingId int, updatedValues *models.Booking) error {
	return c.repo.UpdateBooking(BookingId, updatedValues)
}

func (c *Booking) CancelBooking(BookingId int) error {
	return c.repo.CancelBooking(BookingId)
}
