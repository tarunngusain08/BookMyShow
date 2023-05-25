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

type Theatre struct {
	service services.TheatreServices
}

func NewTheatreController(service services.TheatreServices) *Theatre {
	return &Theatre{
		service: service,
	}
}

func (t *Theatre) GetTheatres(w http.ResponseWriter, req *http.Request) {
	// Extract the city parameter from the query string or request body if needed
	cityId, err := strconv.Atoi(req.FormValue("cityId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.BadRequestError))
		return
	}
	theatres, err := t.service.GetTheatres(cityId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	response, err := json.Marshal(theatres)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.MarshalError))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (t *Theatre) AddTheatre(w http.ResponseWriter, req *http.Request) {
	theatre := &models.Theatre{}
	err := json.NewDecoder(req.Body).Decode(theatre)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.BadRequestError))
		return
	}

	err = t.service.AddTheatre(theatre)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (t *Theatre) GetTheatre(w http.ResponseWriter, req *http.Request) {
	theatreID, err := strconv.Atoi(mux.Vars(req)["theatre_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidTheatreID))
		return
	}

	theatre, err := t.service.GetTheatre(theatreID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	response, err := json.Marshal(theatre)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.MarshalError))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (t *Theatre) UpdateTheatre(w http.ResponseWriter, req *http.Request) {
	theatreID, err := strconv.Atoi(mux.Vars(req)["theatre_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidTheatreID))
		return
	}

	updatedTheatre := &models.Theatre{}
	err = json.NewDecoder(req.Body).Decode(updatedTheatre)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.BadRequestError))
		return
	}

	err = t.service.UpdateTheatre(theatreID, updatedTheatre)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (t *Theatre) DeleteTheatre(w http.ResponseWriter, req *http.Request) {
	theatreID, err := strconv.Atoi(mux.Vars(req)["theatre_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidTheatreID))
		return
	}

	err = t.service.DeleteTheatre(theatreID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}
