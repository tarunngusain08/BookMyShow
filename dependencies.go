package main

import (
	"database/sql"

	"github.com/dunzoit/BookMyShow/controllers"
	"github.com/dunzoit/BookMyShow/repos"
	"github.com/dunzoit/BookMyShow/services"
)

type Controller struct {
	City       *controllers.City
	Theatre    *controllers.Theatre
	Auditorium *controllers.Auditorium
	Show       *controllers.Show
	Movie      *controllers.Movie
	Seat       *controllers.Seat
	Payment    *controllers.Payment
	Booking    *controllers.Booking
	User       *controllers.User
}

func InitializeDependencies() *Controller {

	db := &sql.DB{}
	CityRepo := repos.NewCityRepo(db)
	CityService := services.NewCityService(CityRepo)
	CityController := controllers.NewCityController(CityService)

	TheatreRepo := repos.NewTheatreRepo(db)
	TheatreService := services.NewTheatreService(TheatreRepo)
	TheatreController := controllers.NewTheatreController(TheatreService)

	AuditoriumRepo := repos.NewAuditoriumRepo(db)
	AuditoriumService := services.NewAuditoriumService(AuditoriumRepo)
	AuditoriumController := controllers.NewAuditoriumController(AuditoriumService)

	ShowRepo := repos.NewShowRepo(db)
	ShowService := services.NewShowService(ShowRepo)
	ShowController := controllers.NewShowController(ShowService)

	MovieRepo := repos.NewMovieRepo(db)
	MovieService := services.NewMovieService(MovieRepo)
	MovieController := controllers.NewMovieController(MovieService)

	SeatRepo := repos.NewSeatRepo(db)
	SeatService := services.NewSeatService(SeatRepo)
	SeatController := controllers.NewSeatController(SeatService)

	PaymentRepo := repos.NewPaymentRepo(db)
	PaymentService := services.NewPaymentService(PaymentRepo)
	PaymentController := controllers.NewPaymentController(PaymentService)

	BookingRepo := repos.NewBookingRepo(db)
	BookingService := services.NewBookingService(BookingRepo)
	BookingController := controllers.NewBookingController(BookingService)

	UserRepo := repos.NewUserRepo(db)
	UserService := services.NewUserService(UserRepo)
	UserController := controllers.NewUserController(UserService)

	return &Controller{
		City:       CityController,
		Theatre:    TheatreController,
		Auditorium: AuditoriumController,
		Show:       ShowController,
		Seat:       SeatController,
		Movie:      MovieController,
		Payment:    PaymentController,
		Booking:    BookingController,
		User:       UserController,
	}
}
