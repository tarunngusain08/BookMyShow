package services

import (
	"github.com/dunzoit/BookMyShow/dtos"
	"github.com/dunzoit/BookMyShow/models"
)

type CityServices interface {
	GetCities() ([]*models.City, error)
	AddCities(cities *dtos.CreateCitiesRequest) error
	GetCity(cityId int) (*models.City, error)
	UpdateCity(cityId int, updatedValues *models.City) error
	DeleteCity(cityId int) error
}

type TheatreServices interface {
	GetTheatres(cityID int) ([]*models.Theatre, error)
	AddTheatre(theatre *models.Theatre) error
	GetTheatre(theatreID int) (*models.Theatre, error)
	UpdateTheatre(theatreID int, updatedValues *models.Theatre) error
	DeleteTheatre(theatreID int) error
}

type AuditoriumServices interface {
	GetAuditoriums() ([]*models.Auditorium, error)
	AddAuditorium(auditorium *models.Auditorium) error
	GetAuditorium(AuditoriumId int) (*models.Auditorium, error)
	UpdateAuditorium(AuditoriumId int, updatedValues *models.Auditorium) error
	DeleteAuditorium(AuditoriumId int) error
}

type ShowServices interface {
	GetShows() ([]*models.Show, error)
	AddShow(Show *dtos.CreateShowRequest) error
	GetShow(ShowID int) (*models.Show, error)
	UpdateShow(ShowID int, updatedValues *models.Show) error
	DeleteShow(ShowID int) error
}

type BookingServices interface {
	GetBookings(userId int) ([]*models.Booking, error)
	MakeBooking(Booking *models.Booking) error
	GetBooking(BookingId int) (*models.Booking, error)
	UpdateBooking(BookingId int, updatedValues *models.Booking) error
	CancelBooking(BookingId int) error
}

type MovieServices interface {
	GetMovies() ([]*models.Movie, error)
	AddMovie(Movies *dtos.CreateMovieRequest) error
	GetMovie(MovieID int) (*models.Movie, error)
	UpdateMovie(MovieID int, updatedValues *models.Movie) error
	DeleteMovie(MovieID int) error
}

type PaymentServices interface {
	MakePayment(Payment *models.Payment) error
	GetPayment(PaymentId int) (*models.Payment, error)
	UpdatePayment(PaymentId int, updatedValues *models.Payment) error
	CancelPayment(PaymentId int) error
}

type SeatServices interface {
	GetSeats() ([]*models.Seat, error)
	AddSeat(Seat *models.Seat) error
	GetSeat(SeatID int) (*models.Seat, error)
	UpdateSeat(SeatID int, updatedValues *models.Seat) error
	DeleteSeat(SeatID int) error
}

type UserServices interface {
	GetUsers() ([]*models.User, error)
	CreateUser(user *models.User) error
	GetUser(userID int) (*models.User, error)
	UpdateUser(userID int, updatedUser *models.User) error
	DeleteUser(userID int) error
}
