package service

import "github.com/estefacdcm/reservation/internal/domain/dto"

func (rs *ReservationService) GetReservartionByNumber(reservationNumber int64) ([]dto.ReservationDTO, error) {
	return rs.reservationRepository.FindReservartionByNumber(reservationNumber)
}
