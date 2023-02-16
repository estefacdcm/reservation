package reservation

import (
	"net/http"
	"strings"
	"time"

	"github.com/estefacdcm/reservation/internal/infrastructure/entrypoints/utils"
	"github.com/labstack/echo/v4"
)

const invalidParamCustomerID = "Invalid param customer ID"

func (handler *ReservationHandler) GetReservationByCustomerIDHandler(c echo.Context) error {
	customerID := c.QueryParam("customer_id")

	if customerID == "" {
		return c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(false, time.Now(), strings.Split(invalidParamCustomerID, "; ")))
	}

	reservationByCustomerIDResponse, err := handler.reservationUseCase.GetReservationByCustomerIDUseCase(customerID)
	if err != nil {
		return c.JSON(http.StatusConflict, utils.BuildErrorResponse(false, time.Now(), []string{err.Error()}))
	}

	return c.JSON(http.StatusOK, utils.BuildSuccessfulResponse(true, time.Now(), reservationByCustomerIDResponse))
}
