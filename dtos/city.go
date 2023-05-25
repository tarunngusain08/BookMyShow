package dtos

import "BookMyShow/models"

type CitiesResponse struct {
	Cities []*models.City
}

type CreateCitiesRequest struct {
	Cities []*models.City
}

type CityRequest struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
}
