package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"BookMyShow/exceptions"
	"BookMyShow/models"
	"BookMyShow/services"
	"github.com/gorilla/mux"
)

type Booking struct {
	service services.BookingServices
}

func NewBookingController(service services.BookingServices) *Booking {
	return &Booking{
		service: service,
	}
}

func (b *Booking) GetBookings(w http.ResponseWriter, req *http.Request) {
	userID, err := strconv.Atoi(req.Header.Get("User-ID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.BadRequestError))
	}

	bookings, err := b.service.GetBookings(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	response, err := json.Marshal(bookings)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.MarshalError))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (b *Booking) MakeBooking(w http.ResponseWriter, req *http.Request) {
	booking := &models.Booking{}
	err := json.NewDecoder(req.Body).Decode(booking)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.BadRequestError))
		return
	}

	err = b.service.MakeBooking(booking)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (b *Booking) GetBooking(w http.ResponseWriter, req *http.Request) {
	bookingID, err := strconv.Atoi(mux.Vars(req)["booking_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidBookingID))
		return
	}

	booking, err := b.service.GetBooking(bookingID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	response, err := json.Marshal(booking)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.MarshalError))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (b *Booking) UpdateBooking(w http.ResponseWriter, req *http.Request) {
	bookingID, err := strconv.Atoi(mux.Vars(req)["booking_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidBookingID))
		return
	}

	updatedBooking := &models.Booking{}
	err = json.NewDecoder(req.Body).Decode(updatedBooking)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.BadRequestError))
		return
	}

	err = b.service.UpdateBooking(bookingID, updatedBooking)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (b *Booking) CancelBooking(w http.ResponseWriter, req *http.Request) {
	bookingID, err := strconv.Atoi(mux.Vars(req)["booking_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidBookingID))
		return
	}

	err = b.service.CancelBooking(bookingID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}
