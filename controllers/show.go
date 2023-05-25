package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dunzoit/BookMyShow/dtos"
	"github.com/dunzoit/BookMyShow/exceptions"
	"github.com/dunzoit/BookMyShow/models"
	"github.com/dunzoit/BookMyShow/services"
	"github.com/gorilla/mux"
)

type Show struct {
	service services.ShowServices
}

func NewShowController(service services.ShowServices) *Show {
	return &Show{
		service: service,
	}
}

func (s *Show) GetShows(w http.ResponseWriter, req *http.Request) {
	shows, err := s.service.GetShows()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	response, err := json.Marshal(shows)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.MarshalError))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (s *Show) AddShow(w http.ResponseWriter, req *http.Request) {
	show := &dtos.CreateShowRequest{}
	err := json.NewDecoder(req.Body).Decode(show)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.BadRequestError))
		return
	}

	err = s.service.AddShow(show)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Show) GetShow(w http.ResponseWriter, req *http.Request) {
	showID, err := strconv.Atoi(mux.Vars(req)["show_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidShowID))
		return
	}

	show, err := s.service.GetShow(showID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	response, err := json.Marshal(show)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.MarshalError))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (s *Show) UpdateShow(w http.ResponseWriter, req *http.Request) {
	showID, err := strconv.Atoi(mux.Vars(req)["show_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidShowID))
		return
	}

	updatedShow := &models.Show{}
	err = json.NewDecoder(req.Body).Decode(updatedShow)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.BadRequestError))
		return
	}

	err = s.service.UpdateShow(showID, updatedShow)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Show) DeleteShow(w http.ResponseWriter, req *http.Request) {
	showID, err := strconv.Atoi(mux.Vars(req)["show_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidShowID))
		return
	}

	err = s.service.DeleteShow(showID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}
