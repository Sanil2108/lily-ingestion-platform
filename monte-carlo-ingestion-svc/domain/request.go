package domain

// Request is the Incoming Request for sanitization, validation and handling the request
type Request struct {
	PathParams  interface{}       `json:"params"`
	QueryParams interface{}       `json:"query"`
	Body        interface{}       `json:"body"`
	Headers     interface{}       `json:"header"`
}

type IngestionRequest struct {
	APIKey string `json:"apiKey" validate:"required"`
}

type IngestionResponse struct {
	Status string `json:"status"`
}