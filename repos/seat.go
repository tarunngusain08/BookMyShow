package repos

import (
	"database/sql"

	"BookMyShow/models"
)

const (
	getSeats    = `SELECT id, type, price, availability, auditorium_id FROM seats`
	addSeat     = `INSERT INTO seats (type, price, availability, auditorium_id) VALUES ($1, $2, $3, $4) RETURNING id`
	getSeatByID = `SELECT id, type, price, availability, auditorium_id FROM seats WHERE id = $1`
	updateSeat  = `UPDATE seats SET type = $1, price = $2, availability = $3, auditorium_id = $4 WHERE id = $5`
	deleteSeat  = `DELETE FROM seats WHERE id = $1`
)

type Seat struct {
	db *sql.DB
}

func NewSeatRepo(db *sql.DB) *Seat {
	return &Seat{
		db: db,
	}
}

func (s *Seat) GetSeats() ([]*models.Seat, error) {
	rows, err := s.db.Query(getSeats)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seats []*models.Seat
	for rows.Next() {
		var id int
		var seatType string
		var price float64
		var availability bool
		var auditoriumID int
		if err := rows.Scan(&id, &seatType, &price, &availability, &auditoriumID); err != nil {
			return nil, err
		}
		seat := &models.Seat{
			ID:           id,
			Type:         seatType,
			Price:        price,
			Availability: availability,
			Auditorium:   &models.Auditorium{ID: auditoriumID},
		}
		seats = append(seats, seat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return seats, nil
}

func (s *Seat) AddSeat(seat *models.Seat) error {
	var seatID int
	err := s.db.QueryRow(addSeat, seat.Type, seat.Price, seat.Availability, seat.Auditorium.ID).Scan(&seatID)
	if err != nil {
		return err
	}
	seat.ID = seatID
	return nil
}

func (s *Seat) GetSeatByID(seatID int) (*models.Seat, error) {
	row := s.db.QueryRow(getSeatByID, seatID)

	var id int
	var seatType string
	var price float64
	var availability bool
	var auditoriumID int
	if err := row.Scan(&id, &seatType, &price, &availability, &auditoriumID); err != nil {
		return nil, err
	}

	seat := &models.Seat{
		ID:           id,
		Type:         seatType,
		Price:        price,
		Availability: availability,
		Auditorium:   &models.Auditorium{ID: auditoriumID},
	}

	return seat, nil
}

func (s *Seat) UpdateSeat(seatID int, updatedValues *models.Seat) error {
	_, err := s.db.Exec(updateSeat, updatedValues.Type, updatedValues.Price, updatedValues.Availability,
		updatedValues.Auditorium.ID, seatID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Seat) DeleteSeat(seatID int) error {
	_, err := s.db.Exec(deleteSeat, seatID)
	if err != nil {
		return err
	}
	return nil
}
