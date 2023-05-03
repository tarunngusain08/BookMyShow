package models

import "time"

type Show struct {
	Id        int
	StartTime *time.Time
	EndTime   *time.Time
}
