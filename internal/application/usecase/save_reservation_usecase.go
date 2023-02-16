package usecase

import (
	"time"

	"github.com/estefacdcm/reservation/internal/application/model"
	"github.com/estefacdcm/reservation/internal/domain/dto"
	"github.com/estefacdcm/reservation/internal/domain/entity"
	"github.com/labstack/gommon/log"
)

func (ruc *ReservationUseCase) SaveReservationUseCase(reservationModel *model.ReservationModel) (*dto.ReservationDTO, error) {
	reservationEntity := buildReservationTransaction(reservationModel)

	resultCreate, err := ruc.ReservationService.CreateReservation(reservationEntity)
	if err != nil {
		log.Error("Error creating reservation", err)
		return nil, err
	}

	return &dto.ReservationDTO{
		CustomerID:        resultCreate.CustomerID,
		ReservationNumber: resultCreate.ReservationNumber,
		CreationDate:      resultCreate.CreationDate,
	}, nil
}

func buildReservationTransaction(reservationModel *model.ReservationModel) *entity.ReservationEntity {
	return &entity.ReservationEntity{
		CustomerID:        reservationModel.CustomerID,
		ReservationNumber: reservationModel.ReservationNumber,
		CreationDate:      time.Now(),
	}
}
