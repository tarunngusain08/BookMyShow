package models

type Booking struct {
    ID           int
    User         *User
    Show         *Show
    NumOfSeats   int
    TotalAmount  float64
    Payment      *Payment
    BookedSeats  []*Seat
    BookingTime  int64
    Cancelled    bool
}
