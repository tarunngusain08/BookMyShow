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

}
