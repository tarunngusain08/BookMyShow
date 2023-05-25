package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"BookMyShow/exceptions"
	"BookMyShow/models"
	"BookMyShow/services"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

type Seat struct {
	service services.SeatServices
	redis   *redis.Client
}

func NewSeatController(service services.SeatServices, redis *redis.Client) *Seat {
	return &Seat{
		service: service,
		redis:   redis,
	}
}

func (s *Seat) GetSeats(w http.ResponseWriter, req *http.Request) {
	seats, err := s.service.GetSeats()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	response, err := json.Marshal(seats)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.MarshalError))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (s *Seat) AddSeat(w http.ResponseWriter, req *http.Request) {
	seat := &models.Seat{}
	err := json.NewDecoder(req.Body).Decode(seat)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.BadRequestError))
		return
	}

	err = s.service.AddSeat(seat)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Seat) GetSeat(w http.ResponseWriter, req *http.Request) {
	seatId := mux.Vars(req)["seat_id"]
	val := s.redis.Get(context.Background(), seatId)
	if val.Val() != "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.SeatOccupied))
		return
	}

	seatID, err := strconv.Atoi(seatId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidSeatID))
		return
	}

	seat, err := s.service.GetSeat(seatID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	s.redis.Set(context.Background(), seatId, true, 600)

	response, err := json.Marshal(seat)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.MarshalError))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (s *Seat) UpdateSeat(w http.ResponseWriter, req *http.Request) {
	seatID, err := strconv.Atoi(mux.Vars(req)["seat_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidSeatID))
		return
	}

	updatedSeat := &models.Seat{}
	err = json.NewDecoder(req.Body).Decode(updatedSeat)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.BadRequestError))
		return
	}

	err = s.service.UpdateSeat(seatID, updatedSeat)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Seat) DeleteSeat(w http.ResponseWriter, req *http.Request) {
	seatID, err := strconv.Atoi(mux.Vars(req)["seat_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidSeatID))
		return
	}

	err = s.service.DeleteSeat(seatID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}
