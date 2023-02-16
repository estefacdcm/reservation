package reservation

import "github.com/estefacdcm/reservation/internal/infrastructure/entrypoints/port"

type ReservationHandler struct {
	reservationUseCase port.IReservationUseCase
}

func NewReservationHandler(reservationUseCase port.IReservationUseCase) *ReservationHandler {
	return &ReservationHandler{
		reservationUseCase: reservationUseCase,
	}
}
