package port

import (
	"github.com/estefacdcm/reservation/internal/domain/dto"
	"github.com/estefacdcm/reservation/internal/domain/entity"
)

type IReservationService interface {
	CreateReservation(reservationEntity *entity.ReservationEntity) (*entity.ReservationEntity, error)
	GetReservationByCustomerID(customerID string) ([]dto.ReservationDTO, error)
	GetReservartionByNumber(reservationNumber int64) ([]dto.ReservationDTO, error)
}
