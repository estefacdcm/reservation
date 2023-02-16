package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealthHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/reservation/health", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewHealthHandler()

	if assert.NoError(t, h.Health(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
