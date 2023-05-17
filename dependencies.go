package main

import (
	"database/sql"

	"github.com/dunzoit/BookMyShow/controllers"
	"github.com/dunzoit/BookMyShow/repos"
	"github.com/dunzoit/BookMyShow/services"
)

type Controller struct {
	City *controllers.City
}

func InitializeDependencies() *Controller {

	db := &sql.DB{}
	cityRepo := repos.NewCityRepo(db)
	cityService := services.NewCityService(cityRepo)
	cityController := controllers.NewCityController(cityService)

	return &Controller{
		City: cityController,
	}
}
