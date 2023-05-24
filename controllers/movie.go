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

type Movie struct {
	service services.MovieServices
}

func NewMovieController(service services.MovieServices) *Movie {
	return &Movie{
		service: service,
	}
}

func (m *Movie) GetMovies(w http.ResponseWriter, req *http.Request) {
	movies, err := m.service.GetMovies()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	response, err := json.Marshal(movies)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.MarshalError))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (m *Movie) AddMovie(w http.ResponseWriter, req *http.Request) {
	movie := &dtos.CreateMovieRequest{}
	err := json.NewDecoder(req.Body).Decode(movie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.BadRequestError))
		return
	}

	err = m.service.AddMovie(movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (m *Movie) GetMovie(w http.ResponseWriter, req *http.Request) {
	movieID, err := strconv.Atoi(mux.Vars(req)["movie_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidMovieID))
		return
	}

	movie, err := m.service.GetMovie(movieID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	response, err := json.Marshal(movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.MarshalError))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (m *Movie) UpdateMovie(w http.ResponseWriter, req *http.Request) {
	movieID, err := strconv.Atoi(mux.Vars(req)["movie_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidMovieID))
		return
	}

	updatedMovie := &models.Movie{}
	err = json.NewDecoder(req.Body).Decode(updatedMovie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.BadRequestError))
		return
	}

	err = m.service.UpdateMovie(movieID, updatedMovie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (m *Movie) DeleteMovie(w http.ResponseWriter, req *http.Request) {
	movieID, err := strconv.Atoi(mux.Vars(req)["movie_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidMovieID))
		return
	}

	err = m.service.DeleteMovie(movieID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}
