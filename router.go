package main

import (
	"database/sql"
	"net/http"

	"cities/handler"
	"cities/repo"
	"cities/service"
	"github.com/gorilla/mux"
)

type Dependencies struct {
	Cities *handler.CitiesHandler
}

func main() {
	r := mux.NewRouter()

	handler := initializeDependencies()

	r.HandleFunc("/cities", handler.Cities.CreateCity).Methods(http.MethodPost)

	http.ListenAndServe(":8080", r)
}

func initializeDependencies() *Dependencies {
	db := &sql.DB{}
	citiesRepo := repo.NewCitiesRepo(db)
	citiesService := service.NewCitiesService(citiesRepo)
	citiesHandler := handler.NewCitiesHandler(citiesService)
	return &Dependencies{
		Cities: citiesHandler,
	}
}
