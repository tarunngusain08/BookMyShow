package models

import "time"

type Payment struct {
    ID           int
    PaymentType  string
    Amount       float64
    Status       string
    Transaction  string
    Time         *time.Time
}
