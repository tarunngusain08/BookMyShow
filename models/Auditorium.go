package models

type Auditorium struct {
    ID         int
    Name       string
    Theatre    *Theatre
    Features   []string
    Seats      []*Seat
}
