package main

import (
	"github.com/dunzoit/BookMyShow/controllers"
	"github.com/dunzoit/BookMyShow/repos"
	"github.com/dunzoit/BookMyShow/services"
	"github.com/dunzoit/focus-list-management/pkg/storage"
)

type Controller struct {
	City *controllers.City
}

func InitializeDependencies() *Controller {

	db := storage.DBClient{}
	cityRepo := repos.NewCityRepo(db)
	cityService := services.NewCityService(cityRepo)
	cityController := controllers.NewCityController(cityService)

	return &Controller{
		City: cityController,
	}
}
