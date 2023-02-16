package service

import "github.com/estefacdcm/reservation/internal/domain/dto"

func (rs *ReservationService) GetReservationByCustomerID(customerID string) ([]dto.ReservationDTO, error) {
	return rs.reservationRepository.FindReservationByCustomerID(customerID)
}
