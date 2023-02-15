package utils

import (
	"time"

	"github.com/estefacdcm/reservation/internal/infrastructure/entrypoints/api"
)

func BuildErrorResponse(success bool, timeStamp time.Time, errorSlice []string) api.ResponseError {
	return api.ResponseError{
		Success:   success,
		Timestamp: timeStamp,
		Error:     errorSlice,
	}
}
