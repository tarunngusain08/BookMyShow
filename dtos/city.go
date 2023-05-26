package dtos

import "BookMyShow/models"

type CitiesResponse struct {
	Cities []*models.City `json:"cities"`
}

type AddCityRequest struct {
	City *models.City `json:"city"`
}

type CityRequest struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
}
