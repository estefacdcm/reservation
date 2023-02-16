package entity

import "time"

type ReservationEntity struct {
	ID                int64
	CustomerID        string
	ReservationNumber int64
	CreationDate      time.Time
}
