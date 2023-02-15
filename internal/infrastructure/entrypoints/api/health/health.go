package health

import (
	"net/http"
	"time"

	"github.com/estefacdcm/reservation/internal/infrastructure/entrypoints/api"
	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c echo.Context) error {
	reponse := api.ResponseSuccess{
		Success:   true,
		Data:      "OK",
		Timestamp: time.Now(),
	}
	return c.JSON(http.StatusOK, reponse)
}
