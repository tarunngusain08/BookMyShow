package repos

import (
	"database/sql"
	"time"

	"github.com/dunzoit/BookMyShow/models"
)

const (
	addPayment     = `INSERT INTO payments (payment_type, amount, status, transaction, time) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	getPaymentByID = `SELECT id, payment_type, amount, status, transaction, time FROM payments WHERE id = $1`
	updatePayment  = `UPDATE payments SET payment_type = $1, amount = $2, status = $3, transaction = $4, time = $5 WHERE id = $6`
	deletePayment  = `DELETE FROM payments WHERE id = $1`
)

type Payment struct {
	db *sql.DB
}

func NewPaymentRepo(db *sql.DB) *Payment {
	return &Payment{
		db: db,
	}
}

func (p *Payment) AddPayment(payment *models.Payment) error {
	var paymentID int
	err := p.db.QueryRow(addPayment, payment.PaymentType, payment.Amount, payment.Status, payment.Transaction,
		payment.Time).Scan(&paymentID)
	if err != nil {
		return err
	}
	payment.ID = paymentID
	return nil
}

func (p *Payment) GetPaymentByID(paymentID int) (*models.Payment, error) {
	row := p.db.QueryRow(getPaymentByID, paymentID)

	var id int
	var paymentType string
	var amount float64
	var status string
	var transaction string
	var time *time.Time
	if err := row.Scan(&id, &paymentType, &amount, &status, &transaction, &time); err != nil {
		return nil, err
	}

	payment := &models.Payment{
		ID:          id,
		PaymentType: paymentType,
		Amount:      amount,
		Status:      status,
		Transaction: transaction,
		Time:        time,
	}

	return payment, nil
}

func (p *Payment) UpdatePayment(paymentID int, updatedValues *models.Payment) error {
	_, err := p.db.Exec(updatePayment, updatedValues.PaymentType, updatedValues.Amount, updatedValues.Status,
		updatedValues.Transaction, updatedValues.Time, paymentID)
	if err != nil {
		return err
	}
	return nil
}

func (p *Payment) DeletePayment(paymentID int) error {
	_, err := p.db.Exec(deletePayment, paymentID)
	if err != nil {
		return err
	}
	return nil
}
