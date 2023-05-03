package models

type Booking struct {
	Id           int
	UserId       int
	MovieId      int
	TheatreId    int
	AuditoriumId int
	SeatId       int
	ShowId       int
	PaymentId    int
	Status       string
}
