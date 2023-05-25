package services

import (
	"BookMyShow/dtos"
	"BookMyShow/models"
	"BookMyShow/repos"
)

type City struct {
	repo repos.CityRepository
}

func NewCityService(repo repos.CityRepository) *City {
	return &City{
		repo: repo,
	}
}

func (c *City) GetCities() ([]*models.City, error) {
	return c.repo.GetCities()
}

func (c *City) AddCities(cities *dtos.CreateCitiesRequest) error {
	return c.repo.AddCities(cities)
}

func (c *City) GetCity(cityId int) (*models.City, error) {
	return c.repo.GetCity(cityId)
}

func (c *City) UpdateCity(cityId int, updatedValues *models.City) error {
	return c.repo.UpdateCity(cityId, updatedValues)
}

func (c *City) DeleteCity(cityId int) error {
	return c.repo.DeleteCity(cityId)
}
