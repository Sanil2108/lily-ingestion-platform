package domain

// Request is the Incoming Request for sanitization, validation and handling the request
type Request struct {
	PathParams  interface{} `json:"params"`
	QueryParams interface{} `json:"query"`
	Body        interface{} `json:"body"`
	Headers     interface{} `json:"header"`
}

type IngestionRequestHealth struct {
	Status    string `json:"status" validate:"required"`
	TableName string `json:"tableName" validate:"required"`
	Timestamp int    `json:"timestamp" validate:"required"`
}

type IngestionRequest struct {
	APIKey       string                 `json:"apiKey" validate:"required"`
	UserId       string                 `json:"userId" validate:"required"`
	TenantId     string                 `json:"tenantId" validate:"required"`
	HealthStatus IngestionRequestHealth `json:"healthStatus" validate:"required"`
}

type IngestionResponse struct {
	Status     string `json:"status"`
	WorkflowId string `json:"workflowId"`
}
