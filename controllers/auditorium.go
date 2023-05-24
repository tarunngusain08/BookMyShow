package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dunzoit/BookMyShow/exceptions"
	"github.com/dunzoit/BookMyShow/models"
	"github.com/dunzoit/BookMyShow/services"
	"github.com/gorilla/mux"
)

type Auditorium struct {
	service services.AuditoriumServices
}

func NewAuditoriumController(service services.AuditoriumServices) *Auditorium {
	return &Auditorium{
		service: service,
	}
}

func (a *Auditorium) GetAuditoriums(w http.ResponseWriter, req *http.Request) {

	auditoriums, err := a.service.GetAuditoriums()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	response, err := json.Marshal(auditoriums)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.MarshalError))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (a *Auditorium) AddAuditorium(w http.ResponseWriter, req *http.Request) {
	auditorium := &models.Auditorium{}
	err := json.NewDecoder(req.Body).Decode(auditorium)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.BadRequestError))
		return
	}

	err = a.service.AddAuditorium(auditorium)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *Auditorium) GetAuditorium(w http.ResponseWriter, req *http.Request) {
	auditoriumID, err := strconv.Atoi(mux.Vars(req)["auditorium_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidAuditoriumID))
		return
	}

	auditorium, err := a.service.GetAuditorium(auditoriumID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	response, err := json.Marshal(auditorium)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.MarshalError))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (a *Auditorium) UpdateAuditorium(w http.ResponseWriter, req *http.Request) {
	auditoriumID, err := strconv.Atoi(mux.Vars(req)["auditorium_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidAuditoriumID))
		return
	}

	updatedAuditorium := &models.Auditorium{}
	err = json.NewDecoder(req.Body).Decode(updatedAuditorium)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.BadRequestError))
		return
	}

	err = a.service.UpdateAuditorium(auditoriumID, updatedAuditorium)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *Auditorium) DeleteAuditorium(w http.ResponseWriter, req *http.Request) {
	auditoriumID, err := strconv.Atoi(mux.Vars(req)["auditorium_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidAuditoriumID))
		return
	}

	err = a.service.DeleteAuditorium(auditoriumID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}
