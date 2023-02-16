package model

type ReservationModel struct {
	CustomerID        string `json:"customer_id"`
	ReservationNumber int64  `json:"reservation_number"`
}
