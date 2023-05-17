package dtos

import "github.com/dunzoit/BookMyShow/models"

type GetCitiesResponse struct {
	Cities []*models.City
}

type AddCitiesRequest struct {
	Cities []*models.City
}

type GetCityResponse struct {
	City *models.City
}

type UpdateCityRequest struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
}
