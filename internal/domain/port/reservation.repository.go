package port

import (
	"github.com/estefacdcm/reservation/internal/domain/dto"
	"github.com/estefacdcm/reservation/internal/domain/entity"
)

type IReservationRepository interface {
	SaveReservation(reservationEntity *entity.ReservationEntity) (*entity.ReservationEntity, error)
	FindReservationByCustomerID(customerID string) ([]dto.ReservationDTO, error)
	FindReservartionByNumber(reservationNumber int64) ([]dto.ReservationDTO, error)
}
