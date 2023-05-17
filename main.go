package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	controller := InitializeDependencies()

	router := mux.NewRouter()
	router.HandleFunc("/cities", controller.City.GetCities).Methods(http.MethodGet)
	router.HandleFunc("/cities", controller.City.AddCities).Methods(http.MethodPost)
	router.HandleFunc("/cities/{city-id}", controller.City.GetCity).Methods(http.MethodGet)
	router.HandleFunc("/cities/{city-id}", controller.City.UpdateCity).Methods(http.MethodPut)
	router.HandleFunc("/cities/{city-id}", controller.City.DeleteCity).Methods(http.MethodDelete)

}
