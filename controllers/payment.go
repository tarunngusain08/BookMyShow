package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dunzoit/BookMyShow/exceptions"
	"github.com/dunzoit/BookMyShow/models"
	"github.com/dunzoit/BookMyShow/services"
	"github.com/gorilla/mux"
)

type Payment struct {
	service services.PaymentServices
}

func NewPaymentController(service services.PaymentServices) *Payment {
	return &Payment{
		service: service,
	}
}

func (p *Payment) MakePayment(w http.ResponseWriter, req *http.Request) {
	payment := &models.Payment{}
	err := json.NewDecoder(req.Body).Decode(payment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.BadRequestError))
		return
	}

	err = p.service.MakePayment(payment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (p *Payment) GetPayment(w http.ResponseWriter, req *http.Request) {
	paymentID, err := strconv.Atoi(mux.Vars(req)["payment_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidPaymentID))
		return
	}

	payment, err := p.service.GetPayment(paymentID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	response, err := json.Marshal(payment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.MarshalError))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (p *Payment) UpdatePayment(w http.ResponseWriter, req *http.Request) {
	paymentID, err := strconv.Atoi(mux.Vars(req)["payment_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidPaymentID))
		return
	}

	updatedPayment := &models.Payment{}
	err = json.NewDecoder(req.Body).Decode(updatedPayment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.BadRequestError))
		return
	}

	err = p.service.UpdatePayment(paymentID, updatedPayment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (p *Payment) CancelPayment(w http.ResponseWriter, req *http.Request) {
	paymentID, err := strconv.Atoi(mux.Vars(req)["payment_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(exceptions.InvalidPaymentID))
		return
	}

	err = p.service.CancelPayment(paymentID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(exceptions.InternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
}
