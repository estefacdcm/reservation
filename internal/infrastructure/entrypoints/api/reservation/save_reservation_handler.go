package reservation

import (
	"net/http"
	"time"

	"github.com/estefacdcm/reservation/internal/application/model"
	responseDto "github.com/estefacdcm/reservation/internal/infrastructure/entrypoints/api"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const errorDecodingBody = "Error decoding body"

func (handler *ReservationHandler) SaveReservationHandler(c echo.Context) error {
	var body *model.ReservationModel
	errBind := c.Bind(&body)
	if errBind != nil {
		logrus.Errorf(errorDecodingBody+": %v", errBind)
		response := responseDto.ResponseError{
			Success:   false,
			Error:     []string{errorDecodingBody},
			Timestamp: time.Now(),
			Data:      errBind,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	data, err := handler.reservationUseCase.SaveReservationUseCase(body)
	if err != nil {
		response := responseDto.ResponseError{
			Success:   false,
			Error:     []string{err.Error()},
			Timestamp: time.Now(),
			Data:      body,
		}
		return c.JSON(http.StatusConflict, response)
	}
	response := responseDto.ResponseSuccess{
		Success:   true,
		Data:      data,
		Timestamp: time.Now(),
	}
	return c.JSON(http.StatusCreated, response)
}
