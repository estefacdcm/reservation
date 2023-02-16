package usecase

import (
	"github.com/estefacdcm/reservation/internal/domain/dto"
	"github.com/labstack/gommon/log"
)

func (ruc *ReservationUseCase) GetReservartionByNumberUseCase(reservationNumber int64) ([]dto.ReservationDTO, error) {
	resultGetReservation, err := ruc.ReservationService.GetReservartionByNumber(reservationNumber)
	if err != nil {
		log.Error("Error querying reservation by reservation number", err)
		return nil, err
	}
	return resultGetReservation, nil
}
