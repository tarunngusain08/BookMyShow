package dtos

import "time"

type CreatePaymentRequest struct {
	PaymentType string     `json:"payment_type"`
	Amount      float64    `json:"amount"`
	Status      string     `json:"status"`
	Transaction string     `json:"transaction"`
	Time        *time.Time `json:"time"`
}

type PaymentResponse struct {
	ID          int        `json:"id"`
	PaymentType string     `json:"payment_type"`
	Amount      float64    `json:"amount"`
	Status      string     `json:"status"`
	Transaction string     `json:"transaction"`
	Time        *time.Time `json:"time"`
}
