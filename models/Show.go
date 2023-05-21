package models

import "time"

type Show struct {
	ID         int
	StartTime  *time.Time
	EndTime    *time.Time
	Theatre    *Theatre
	Screen     int
	Movie      *Movie
	Features   []string
	Seats      []*Seat
	BookedSeat map[*Seat]int
}
