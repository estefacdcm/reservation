package service

import "github.com/estefacdcm/reservation/internal/domain/port"

type ReservationService struct {
	reservationRepository port.IReservationRepository
}

func NewReservationService(reservationRepository port.IReservationRepository) *ReservationService {
	return &ReservationService{
		reservationRepository,
	}
}
