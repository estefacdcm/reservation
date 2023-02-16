package usecase

import (
	"github.com/estefacdcm/reservation/internal/application/port"
)

type ReservationUseCase struct {
	ReservationService port.IReservationService
}

func NewReservationUseCase(reservationService port.IReservationService) *ReservationUseCase {
	return &ReservationUseCase{
		ReservationService: reservationService,
	}
}
