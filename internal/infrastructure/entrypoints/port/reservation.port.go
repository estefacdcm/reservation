package port

import (
	"github.com/estefacdcm/reservation/internal/application/model"
	"github.com/estefacdcm/reservation/internal/domain/dto"
)

type IReservationUseCase interface {
	SaveReservationUseCase(reservationModel *model.ReservationModel) (*dto.ReservationDTO, error)
	GetReservationByCustomerIDUseCase(customerID string) ([]dto.ReservationDTO, error)
	GetReservartionByNumberUseCase(reservationNumber int64) ([]dto.ReservationDTO, error)
}
