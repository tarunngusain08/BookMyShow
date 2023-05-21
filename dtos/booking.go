package dtos

type CreateBookingRequest struct {
	UserID      int     `json:"user_id"`
	ShowID      int     `json:"show_id"`
	NumOfSeats  int     `json:"num_of_seats"`
	TotalAmount float64 `json:"total_amount"`
	PaymentID   int     `json:"payment_id"`
	BookedSeats []int   `json:"booked_seats"`
	BookingTime int64   `json:"booking_time"`
	Cancelled   bool    `json:"cancelled"`
}

type BookingResponse struct {
	ID          int              `json:"id"`
	User        *UserResponse    `json:"user"`
	Show        *ShowResponse    `json:"show"`
	NumOfSeats  int              `json:"num_of_seats"`
	TotalAmount float64          `json:"total_amount"`
	Payment     *PaymentResponse `json:"payment"`
	BookedSeats []*SeatResponse  `json:"booked_seats"`
	BookingTime int64            `json:"booking_time"`
	Cancelled   bool             `json:"cancelled"`
}
