package repos

import (
	"database/sql"

	"github.com/dunzoit/BookMyShow/models"
)

const (
	getBookings        = `SELECT id, user_id, show_id, num_of_seats, total_amount, payment_id, booking_time, cancelled FROM bookings WHERE user_id = $1`
	addBooking         = `INSERT INTO bookings (user_id, show_id, num_of_seats, total_amount, payment_id, booking_time, cancelled) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	getBookingByID     = `SELECT id, user_id, show_id, num_of_seats, total_amount, payment_id, booking_time, cancelled FROM bookings WHERE id = $1`
	updateBooking      = `UPDATE bookings SET num_of_seats = $1, total_amount = $2, payment_id = $3, booking_time = $4, cancelled = $5 WHERE id = $6`
	deleteBooking      = `DELETE FROM bookings WHERE id = $1`
	getBookingSeats    = `SELECT seat_id FROM booking_seats WHERE booking_id = $1`
	addBookingSeat     = `INSERT INTO booking_seats (booking_id, seat_id) VALUES ($1, $2)`
	deleteBookingSeats = `DELETE FROM booking_seats WHERE booking_id = $1`
)

type Booking struct {
	db *sql.DB
}

func NewBookingRepo(db *sql.DB) *Booking {
	return &Booking{
		db: db,
	}
}

func (b *Booking) GetBookings(userID int) ([]*models.Booking, error) {
	rows, err := b.db.Query(getBookings, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []*models.Booking
	for rows.Next() {
		var id int
		var userID int
		var showID int
		var numOfSeats int
		var totalAmount float64
		var paymentID int
		var bookingTime int64
		var cancelled bool
		if err := rows.Scan(&id, &userID, &showID, &numOfSeats, &totalAmount, &paymentID, &bookingTime,
			&cancelled); err != nil {
			return nil, err
		}
		booking := &models.Booking{
			ID:          id,
			User:        nil, // You need to populate the user object
			Show:        nil, // You need to populate the show object
			NumOfSeats:  numOfSeats,
			TotalAmount: totalAmount,
			Payment:     nil, // You need to populate the payment object
			BookedSeats: nil, // You need to populate the booked seats slice
			BookingTime: bookingTime,
			Cancelled:   cancelled,
		}
		bookings = append(bookings, booking)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return bookings, nil
}

func (b *Booking) AddBooking(booking *models.Booking) error {
	var bookingID int
	err := b.db.QueryRow(addBooking, booking.User.ID, booking.Show.ID, booking.NumOfSeats, booking.TotalAmount,
		booking.Payment.ID, booking.BookingTime, booking.Cancelled).Scan(&bookingID)
	if err != nil {
		return err
	}
	booking.ID = bookingID

	// Add booked seats
	for _, seat := range booking.BookedSeats {
		_, err := b.db.Exec(addBookingSeat, bookingID, seat.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Booking) GetBookingByID(bookingID int) (*models.Booking, error) {
	row := b.db.QueryRow(getBookingByID, bookingID)

	var id int
	var userID int
	var showID int
	var numOfSeats int
	var totalAmount float64
	var paymentID int
	var bookingTime int64
	var cancelled bool
	if err := row.Scan(&id, &userID, &showID, &numOfSeats, &totalAmount, &paymentID, &bookingTime,
		&cancelled); err != nil {
		return nil, err
	}

	// Populate user object
	userRepo := NewUserRepo(b.db)
	user, err := userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	// Populate show object
	showRepo := NewShowRepo(b.db)
	show, err := showRepo.GetShowByID(showID)
	if err != nil {
		return nil, err
	}

	// Populate payment object
	paymentRepo := NewPaymentRepo(b.db)
	payment, err := paymentRepo.GetPaymentByID(paymentID)
	if err != nil {
		return nil, err
	}

	// Get booked seats
	rows, err := b.db.Query(getBookingSeats, bookingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookedSeats []*models.Seat
	for rows.Next() {
		var seatID int
		if err := rows.Scan(&seatID); err != nil {
			return nil, err
		}

		// Populate seat object
		seatRepo := NewSeatRepo(b.db)
		seat, err := seatRepo.GetSeatByID(seatID)
		if err != nil {
			return nil, err
		}

		bookedSeats = append(bookedSeats, seat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	booking := &models.Booking{
		ID:          id,
		User:        user,
		Show:        show,
		NumOfSeats:  numOfSeats,
		TotalAmount: totalAmount,
		Payment:     payment,
		BookedSeats: bookedSeats,
		BookingTime: bookingTime,
		Cancelled:   cancelled,
	}

	return booking, nil
}

func (b *Booking) UpdateBooking(bookingID int, updatedValues *models.Booking) error {
	_, err := b.db.Exec(updateBooking, updatedValues.NumOfSeats, updatedValues.TotalAmount, updatedValues.Payment.ID,
		updatedValues.BookingTime, updatedValues.Cancelled, bookingID)
	if err != nil {
		return err
	}

	// Delete existing booked seats
	_, err = b.db.Exec(deleteBookingSeats, bookingID)
	if err != nil {
		return err
	}

	// Add updated booked seats
	for _, seat := range updatedValues.BookedSeats {
		_, err := b.db.Exec(addBookingSeat, bookingID, seat.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Booking) DeleteBooking(bookingID int) error {
	_, err := b.db.Exec(deleteBooking, bookingID)
	if err != nil {
		return err
	}

	_, err = b.db.Exec(deleteBookingSeats, bookingID)
	if err != nil {
		return err
	}

	return nil
}
