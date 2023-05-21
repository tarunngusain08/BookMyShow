package dtos

type CreateAuditoriumRequest struct {
	Name      string   `json:"name"`
	TheatreID int      `json:"theatre_id"`
	Features  []string `json:"features"`
	Seats     []int    `json:"seats"`
}

type AuditoriumResponse struct {
	ID       int              `json:"id"`
	Name     string           `json:"name"`
	Theatre  *TheatreResponse `json:"theatre"`
	Features []string         `json:"features"`
	Seats    []*SeatResponse  `json:"seats"`
}
