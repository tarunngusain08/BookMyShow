package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	controller := InitializeDependencies()

	router := mux.NewRouter()
	router.HandleFunc("/cities", controller.City.GetCities).Methods(http.MethodGet)
	router.HandleFunc("/cities", controller.City.AddCity).Methods(http.MethodPost)
	router.HandleFunc("/cities/{city-id}", controller.City.GetCity).Methods(http.MethodGet)
	router.HandleFunc("/cities/{city-id}", controller.City.UpdateCity).Methods(http.MethodPut)
	router.HandleFunc("/cities/{city-id}", controller.City.DeleteCity).Methods(http.MethodDelete)

	router.HandleFunc("/theatres", controller.Theatre.GetTheatres).Methods(http.MethodGet)
	router.HandleFunc("/theatres", controller.Theatre.AddTheatre).Methods(http.MethodPost)
	router.HandleFunc("/theatres/{theatre-id}", controller.Theatre.GetTheatre).Methods(http.MethodGet)
	router.HandleFunc("/theatres/{theatre-id}", controller.Theatre.UpdateTheatre).Methods(http.MethodPut)
	router.HandleFunc("/theatres/{theatre-id}", controller.Theatre.DeleteTheatre).Methods(http.MethodDelete)

	router.HandleFunc("/auditoriums", controller.Auditorium.GetAuditoriums).Methods(http.MethodGet)
	router.HandleFunc("/auditoriums", controller.Auditorium.AddAuditorium).Methods(http.MethodPost)
	router.HandleFunc("/auditoriums/{auditorium-id}", controller.Auditorium.GetAuditorium).Methods(http.MethodGet)
	router.HandleFunc("/auditoriums/{auditorium-id}", controller.Auditorium.UpdateAuditorium).Methods(http.MethodPut)
	router.HandleFunc("/auditoriums/{auditorium-id}", controller.Auditorium.DeleteAuditorium).Methods(http.MethodDelete)

	router.HandleFunc("/Shows", controller.Show.GetShows).Methods(http.MethodGet)
	router.HandleFunc("/Shows", controller.Show.AddShow).Methods(http.MethodPost)
	router.HandleFunc("/Shows/{Show-id}", controller.Show.GetShow).Methods(http.MethodGet)
	router.HandleFunc("/Shows/{Show-id}", controller.Show.UpdateShow).Methods(http.MethodPut)
	router.HandleFunc("/Shows/{Show-id}", controller.Show.DeleteShow).Methods(http.MethodDelete)

	router.HandleFunc("/Movies", controller.Movie.GetMovies).Methods(http.MethodGet)
	router.HandleFunc("/Movies", controller.Movie.AddMovie).Methods(http.MethodPost)
	router.HandleFunc("/Movies/{Movie-id}", controller.Movie.GetMovie).Methods(http.MethodGet)
	router.HandleFunc("/Movies/{Movie-id}", controller.Movie.UpdateMovie).Methods(http.MethodPut)
	router.HandleFunc("/Movies/{Movie-id}", controller.Movie.DeleteMovie).Methods(http.MethodDelete)

	router.HandleFunc("/Seats", controller.Seat.GetSeats).Methods(http.MethodGet)
	router.HandleFunc("/Seats", controller.Seat.AddSeat).Methods(http.MethodPost)
	router.HandleFunc("/Seats/{Seat-id}", controller.Seat.GetSeat).Methods(http.MethodGet)
	router.HandleFunc("/Seats/{Seat-id}", controller.Seat.UpdateSeat).Methods(http.MethodPut)
	router.HandleFunc("/Seats/{Seat-id}", controller.Seat.DeleteSeat).Methods(http.MethodDelete)

	router.HandleFunc("/Payments", controller.Payment.MakePayment).Methods(http.MethodPost)
	router.HandleFunc("/Payments/{Payment-id}", controller.Payment.GetPayment).Methods(http.MethodGet)
	router.HandleFunc("/Payments/{Payment-id}", controller.Payment.UpdatePayment).Methods(http.MethodPut)
	router.HandleFunc("/Payments/{Payment-id}", controller.Payment.CancelPayment).Methods(http.MethodDelete)

	router.HandleFunc("/Bookings", controller.Booking.GetBookings).Methods(http.MethodGet)
	router.HandleFunc("/Bookings", controller.Booking.MakeBooking).Methods(http.MethodPost)
	router.HandleFunc("/Bookings/{Booking-id}", controller.Booking.GetBooking).Methods(http.MethodGet)
	router.HandleFunc("/Bookings/{Booking-id}", controller.Booking.UpdateBooking).Methods(http.MethodPut)
	router.HandleFunc("/Bookings/{Booking-id}", controller.Booking.CancelBooking).Methods(http.MethodDelete)

	router.HandleFunc("/Users", controller.User.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/Users", controller.User.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/Users/{User-id}", controller.User.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/Users/{User-id}", controller.User.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/Users/{User-id}", controller.User.DeleteUser).Methods(http.MethodDelete)

	http.ListenAndServe("localhost:8080", router)
}
