package domain

import "time"

// Response to be returned for the request
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Time    time.Time   `json:"time"`
}

// NewOkResponse is a helper method to create an OK response */
func NewOkResponse(data interface{}) Response {
	return Response{
		Data:    data,
		Time:    time.Now().UTC(),
		Success: true,
	}
}

// NewErrorResponse is a helper method to create a failed Response */
func NewErrorResponse(data interface{}) Response {
	return Response{
		Error:   data,
		Time:    time.Now().UTC(),
		Success: false,
	}
}
