package utils

import (
	"time"

	"github.com/estefacdcm/reservation/internal/infrastructure/entrypoints/api"
)

func BuildSuccessfulResponse(success bool, timeStamp time.Time, data any) api.ResponseSuccess {
	return api.ResponseSuccess{
		Success:   success,
		Data:      data,
		Timestamp: timeStamp,
	}
}
