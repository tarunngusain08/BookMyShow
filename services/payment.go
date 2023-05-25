package services

import (
	"BookMyShow/models"
	"BookMyShow/repos"
)

type Payment struct {
	repo repos.PaymentRepository
}

func NewPaymentService(repo repos.PaymentRepository) *Payment {
	return &Payment{
		repo: repo,
	}
}

func (c *Payment) MakePayment(Payment *models.Payment) error {
	return c.repo.MakePayment(Payment)
}

func (c *Payment) GetPayment(PaymentId int) (*models.Payment, error) {
	return c.repo.GetPaymentByID(PaymentId)
}

func (c *Payment) UpdatePayment(PaymentId int, updatedValues *models.Payment) error {
	return c.repo.UpdatePayment(PaymentId, updatedValues)
}

func (c *Payment) CancelPayment(PaymentId int) error {
	return c.repo.CancelPayment(PaymentId)
}
