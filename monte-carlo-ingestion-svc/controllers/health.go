package controllers

import (
	"context"

	"monte-carlo-ingestion/domain"
)

// HealthController is the controller to handle healthchecks
type HealthController struct {
	DefaultController
}

// Handler handles the health check requests
func (meta *HealthController) Handler(_ context.Context, _ *domain.Request) (interface{}, error) {
	return map[string]string{"status": "ok"}, nil
}

// NewHealthController is the constructor for the HealthController
func NewHealthController() HealthController {
	return HealthController{}
}
