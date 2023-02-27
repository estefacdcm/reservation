package service

import (
	"testing"

	"github.com/estefacdcm/reservation/internal/domain/entity"
	"github.com/estefacdcm/reservation/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateReservation(t *testing.T) {
	t.Run("Should get success response", func(t *testing.T) {
		reservationEntity := entity.ReservationEntity{
			ID:                10,
			CustomerID:        "customer20",
			ReservationNumber: 20,
		}

		reservationRepositoryMock := new(mocks.IReservationRepository)
		reservationRepositoryMock.On("SaveReservation", reservationEntity).Return(&reservationEntity, nil)
		reservationService := NewReservationService(reservationRepositoryMock)
		response, err := reservationService.CreateReservation(&reservationEntity)

		assert.Nil(t, err)
		assert.NotNil(t, response)
	})

}
