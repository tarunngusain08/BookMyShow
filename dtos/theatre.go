package dtos

type TheatreRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	CityID  int    `json:"city_id"`
}

type TheatreResponse struct {
	ID      int           `json:"id"`
	Name    string        `json:"name"`
	Address string        `json:"address"`
	City    *CityResponse `json:"city"`
}

type CityResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
}
