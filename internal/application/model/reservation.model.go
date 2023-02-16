package model

type ReservationModel struct {
	CustomerID        int64 `json:"customer_id"`
	ReservationNumber int64 `json:"reservation_number"`
}
