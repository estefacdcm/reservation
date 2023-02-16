package service

import (
	"testing"

	"github.com/estefacdcm/reservation/mocks"
	"github.com/stretchr/testify/assert"

	"github.com/estefacdcm/reservation/internal/domain/dto"
)

func TestGetReservationByCustomerID(t *testing.T) {
	t.Run("Should get success response", func(t *testing.T) {
		var customerID = "customer20"

		var data []dto.ReservationDTO

		reservation := dto.ReservationDTO{
			ID:                4,
			CustomerID:        "customer14",
			ReservationNumber: 14,
		}
		result := append(data, reservation)

		reservationRepositoryMock := new(mocks.IReservationRepository)
		reservationRepositoryMock.On("FindReservationByCustomerID", customerID).Return(result, nil)
		reservationService := NewReservationService(reservationRepositoryMock)
		response, err := reservationService.GetReservationByCustomerID(customerID)

		assert.Nil(t, err)
		assert.NotNil(t, response)
	})
}
