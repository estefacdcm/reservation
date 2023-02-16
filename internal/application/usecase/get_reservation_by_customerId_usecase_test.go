package usecase

import (
	"errors"
	"testing"

	"github.com/estefacdcm/reservation/mocks"
	"github.com/stretchr/testify/assert"

	"github.com/estefacdcm/reservation/internal/domain/dto"
)

func TestGetReservationByCustomerIDUseCase(t *testing.T) {
	t.Run("Should get successful response", func(t *testing.T) {

		const customerID = "customer14"

		var data []dto.ReservationDTO

		reservation := dto.ReservationDTO{
			ID:                4,
			CustomerID:        "customer14",
			ReservationNumber: 14,
		}
		result := append(data, reservation)

		reservationServiceMock := new(mocks.IReservationService)
		reservationServiceMock.On("GetReservationByCustomerID", customerID).Return(result, nil)
		reservationUseCase := NewReservationUseCase(reservationServiceMock)
		response, err := reservationUseCase.GetReservationByCustomerIDUseCase(customerID)

		assert.NotNil(t, response)
		assert.Nil(t, err)
		assert.Equal(t, result, response)
	})

	t.Run("Should get error", func(t *testing.T) {
		const customerID = "customer30"

		const msnError = "an error occurred while trying to query reservation by customerI"

		reservationServiceMock := new(mocks.IReservationService)
		reservationServiceMock.On("GetReservationByCustomerID", customerID).Return(nil, errors.New(msnError))
		reservationUseCase := NewReservationUseCase(reservationServiceMock)
		response, err := reservationUseCase.GetReservationByCustomerIDUseCase(customerID)

		assert.Nil(t, response)
		assert.NotNil(t, err)
		assert.Equal(t, msnError, err.Error())
	})
}
