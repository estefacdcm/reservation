package api

import "time"

type (
	ResponseSuccess struct {
		Success   bool      `json:"success"`
		Data      any       `json:"data"`
		Timestamp time.Time `json:"timestamp"`
		Metadata  any       `json:"metadata,omitempty"`
	}

	ResponseError struct {
		Success   bool      `json:"success"`
		Error     []string  `json:"error"`
		Timestamp time.Time `json:"timestamp"`
	}
)
