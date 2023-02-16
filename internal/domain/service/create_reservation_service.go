package service

import "github.com/estefacdcm/reservation/internal/domain/entity"

func (rs *ReservationService) CreateReservation(reservationEntity *entity.ReservationEntity) (*entity.ReservationEntity, error) {
	return rs.reservationRepository.SaveReservation(reservationEntity)
}
