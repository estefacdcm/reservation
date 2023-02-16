package service

import (
	"testing"

	"github.com/estefacdcm/reservation/mocks"
	"github.com/stretchr/testify/assert"

	"github.com/estefacdcm/reservation/internal/domain/dto"
)

func TestGetReservartionByNumber(t *testing.T) {
	t.Run("Should get success response", func(t *testing.T) {
		const reservationNumber = int64(14)

		var data []dto.ReservationDTO

		reservation := dto.ReservationDTO{
			ID:                4,
			CustomerID:        "customer14",
			ReservationNumber: 14,
		}
		result := append(data, reservation)

		reservationRepositoryMock := new(mocks.IReservationRepository)
		reservationRepositoryMock.On("FindReservartionByNumber", reservationNumber).Return(result, nil)
		reservationService := NewReservationService(reservationRepositoryMock)
		response, err := reservationService.GetReservartionByNumber(reservationNumber)

		assert.Nil(t, err)
		assert.NotNil(t, response)
	})
}
