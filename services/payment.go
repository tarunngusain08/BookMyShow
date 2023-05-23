package services

import (
	"github.com/dunzoit/BookMyShow/models"
	"github.com/dunzoit/BookMyShow/repos"
)

type Payment struct {
	repo repos.PaymentRepository
}

func NewPaymentService(repo repos.PaymentRepository) *Payment {
	return &Payment{
		repo: repo,
	}
}

func (c *Payment) AddPayment(Payment *models.Payment) error {
	return c.repo.AddPayment(Payment)
}

func (c *Payment) GetPayment(PaymentId int) (*models.Payment, error) {
	return c.repo.GetPaymentByID(PaymentId)
}

func (c *Payment) UpdatePayment(PaymentId int, updatedValues *models.Payment) error {
	return c.repo.UpdatePayment(PaymentId, updatedValues)
}

func (c *Payment) DeletePayment(PaymentId int) error {
	return c.repo.DeletePayment(PaymentId)
}
