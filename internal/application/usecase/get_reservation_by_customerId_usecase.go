package usecase

import (
	"github.com/estefacdcm/reservation/internal/domain/dto"
	"github.com/labstack/gommon/log"
)

func (ruc *ReservationUseCase) GetReservationByCustomerIDUseCase(customerID string) ([]dto.ReservationDTO, error) {
	resultGetReservation, err := ruc.ReservationService.GetReservationByCustomerID(customerID)
	if err != nil {
		log.Error("Error querying reservation by customerID", err)
		return nil, err
	}
	return resultGetReservation, nil
}
