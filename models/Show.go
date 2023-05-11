package models

type Show struct {
    ID         int
    TimeSlot   string
    Theatre    *Theatre
    Screen     int
    Movie      *Movie
    Features   []string
    Seats      []*Seat
    BookedSeat map[*Seat]int
}
