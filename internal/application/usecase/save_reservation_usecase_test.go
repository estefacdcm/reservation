package usecase

import (
	"errors"
	"testing"

	"github.com/estefacdcm/reservation/mocks"
	"github.com/stretchr/testify/assert"

	"github.com/estefacdcm/reservation/internal/application/model"
	"github.com/estefacdcm/reservation/internal/domain/dto"
	"github.com/estefacdcm/reservation/internal/domain/entity"
)

func TestSaveReservationUseCase(t *testing.T) {
	t.Run("Should get successful response", func(t *testing.T) {
		data := dto.ReservationDTO{
			CustomerID:        "curtomer20",
			ReservationNumber: 20,
		}

		reservationModel := model.ReservationModel{
			CustomerID:        "customer20",
			ReservationNumber: 20,
		}

		reservationEntity := entity.ReservationEntity{
			ID:                10,
			CustomerID:        "customer20",
			ReservationNumber: 20,
		}

		reservationServiceMock := new(mocks.IReservationService)
		reservationServiceMock.On("CreateReservation", reservationEntity).Return(data, nil)
		reservationUseCase := NewReservationUseCase(reservationServiceMock)
		response, err := reservationUseCase.SaveReservationUseCase(&reservationModel)

		assert.NotNil(t, response)
		assert.Nil(t, err)
		assert.Equal(t, data, response)
	})

	t.Run("Should get error", func(t *testing.T) {

		reservationModel := model.ReservationModel{
			CustomerID:        "customer20",
			ReservationNumber: 20,
		}

		reservationEntity := entity.ReservationEntity{
			ID:                10,
			CustomerID:        "customer20",
			ReservationNumber: 20,
		}

		const msnError = "error creating reservation"

		reservationServiceMock := new(mocks.IReservationService)
		reservationServiceMock.On("CreateReservation", reservationEntity).Return(nil, errors.New(msnError))
		reservationUseCase := NewReservationUseCase(reservationServiceMock)
		response, err := reservationUseCase.SaveReservationUseCase(&reservationModel)

		assert.Nil(t, response)
		assert.NotNil(t, err)
		assert.Equal(t, msnError, err.Error())
	})
}
