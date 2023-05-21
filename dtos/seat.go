package dtos

type CreateSeatRequest struct {
	Type         string  `json:"type"`
	Price        float64 `json:"price"`
	Availability bool    `json:"availability"`
	AuditoriumID int     `json:"auditorium_id"`
}

type SeatResponse struct {
	ID           int                 `json:"id"`
	Type         string              `json:"type"`
	Price        float64             `json:"price"`
	Availability bool                `json:"availability"`
	Auditorium   *AuditoriumResponse `json:"auditorium"`
}
