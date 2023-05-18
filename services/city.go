package services

import (
	"github.com/dunzoit/BookMyShow/dtos"
	"github.com/dunzoit/BookMyShow/repos"
)

type City struct {
	repo *repos.City
}

func NewCityService(repo *repos.City) *City {
	return &City{
		repo: repo,
	}
}

func (c *City) GetCities() (*dtos.GetCitiesResponse, error) {
	return c.repo.GetCities()
}

func (c *City) AddCities(cities *dtos.AddCitiesRequest) error {
	return c.repo.AddCities(cities)
}

func (c *City) GetCity(cityId int) (*dtos.GetCityResponse, error) {
	return c.repo.GetCity(cityId)
}

func (c *City) UpdateCity(cityId int, updatedValues *dtos.UpdateCityRequest) error {
	return c.repo.UpdateCity(cityId, updatedValues)
}

func (c *City) DeleteCity(cityId int) error {
	return c.repo.DeleteCity(cityId)
}
