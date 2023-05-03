package models

import "time"

type Payment struct {
	Id     int
	UserId int
	Amount int
	Mop    string
	PaidOn *time.Time
}
