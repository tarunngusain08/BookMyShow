package repos

import (
	"github.com/dunzoit/BookMyShow/dtos"
	"github.com/dunzoit/BookMyShow/models"
)

type CityRepository interface {
	GetCities() ([]*models.City, error)
	AddCities(cities *dtos.CreateCitiesRequest) error
	GetCity(cityId int) (*models.City, error)
	UpdateCity(cityId int, updatedValues *models.City) error
	DeleteCity(cityId int) error
}

type TheatreRepository interface {
	GetTheatres(cityID int) ([]*models.Theatre, error)
	AddTheatre(theatre *models.Theatre) error
	GetTheatre(theatreID int) (*models.Theatre, error)
	UpdateTheatre(theatreID int, updatedValues *models.Theatre) error
	DeleteTheatre(theatreID int) error
}

type ShowRepository interface {
	GetShows() ([]*models.Show, error)
	AddShow(show *dtos.CreateShowRequest) error
	GetShowByID(showID int) (*models.Show, error)
	UpdateShow(showID int, updatedValues *models.Show) error
	DeleteShow(showID int) error
}

type MovieRepository interface {
	GetMovies() ([]*models.Movie, error)
	AddMovie(movie *dtos.CreateMovieRequest) error
	GetMovieByID(movieID int) (*models.Movie, error)
	UpdateMovie(movieID int, updatedValues *models.Movie) error
	DeleteMovie(movieID int) error
}

type AuditoriumRepository interface {
	GetAuditoriums() ([]*models.Auditorium, error)
	AddAuditorium(auditorium *models.Auditorium) error
	GetAuditorium(auditoriumID int) (*models.Auditorium, error)
	UpdateAuditorium(auditoriumID int, updatedValues *models.Auditorium) error
	DeleteAuditorium(auditoriumID int) error
}

type SeatRepository interface {
	GetSeats() ([]*models.Seat, error)
	AddSeat(seat *models.Seat) error
	GetSeatByID(seatID int) (*models.Seat, error)
	UpdateSeat(seatID int, updatedValues *models.Seat) error
	DeleteSeat(seatID int) error
}

type PaymentRepository interface {
	MakePayment(payment *models.Payment) error
	GetPaymentByID(paymentID int) (*models.Payment, error)
	UpdatePayment(paymentID int, updatedValues *models.Payment) error
	CancelPayment(paymentID int) error
}

type BookingRepository interface {
	GetBookings(userID int) ([]*models.Booking, error)
	MakeBooking(booking *models.Booking) error
	GetBookingByID(bookingID int) (*models.Booking, error)
	UpdateBooking(bookingID int, updatedValues *models.Booking) error
	CancelBooking(bookingID int) error
}

type UserRepository interface {
	GetUsers() ([]*models.User, error)
	CreateUser(user *models.User) error
	GetUserByID(userID int) (*models.User, error)
	UpdateUser(userID int, updatedValues *models.User) error
	DeleteUser(userID int) error
}
