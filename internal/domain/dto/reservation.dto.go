package dto

import "time"

type ReservationDTO struct {
	ID                int64
	CustomerID        string
	ReservationNumber int64
	CreationDate      time.Time
}
