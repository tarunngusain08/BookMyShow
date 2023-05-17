package models

type Seat struct {
    ID           int
    Type         string
    Price        float64
    Availability bool
    Auditorium   *Auditorium
}
