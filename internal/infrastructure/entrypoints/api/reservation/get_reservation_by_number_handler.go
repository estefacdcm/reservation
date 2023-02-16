package reservation

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/estefacdcm/reservation/internal/infrastructure/entrypoints/utils"
	"github.com/labstack/echo/v4"
)

const (
	invalidParamReservationNumber = "Invalid param Reservation Number"
	errorGettingResrvationNumber  = "Error getting reservation number"
)

func (handler *ReservationHandler) GetReservationByNumberHandler(c echo.Context) error {
	reservationNumberQP := c.QueryParam("reservation_number")

	if reservationNumberQP != "" {
		return c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(false, time.Now(), strings.Split(invalidParamReservationNumber, "; ")))
	}

	reservationNumber, err := strconv.ParseInt(reservationNumberQP, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(false, time.Now(), strings.Split(errorGettingResrvationNumber, "; ")))
	}
	reservationByNumberResponse, err := handler.reservationUseCase.GetReservartionByNumberUseCase(int64(reservationNumber))
	if err != nil {
		return c.JSON(http.StatusConflict, utils.BuildErrorResponse(false, time.Now(), []string{err.Error()}))
	}

	return c.JSON(http.StatusOK, utils.BuildSuccessfulResponse(true, time.Now(), reservationByNumberResponse))
}
