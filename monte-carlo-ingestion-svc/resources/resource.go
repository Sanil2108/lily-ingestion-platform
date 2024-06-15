package resources

import (
	"context"

	"monte-carlo-ingestion/config"
)

// Resource is the base resource interface
type Resource interface {
	GetIdentifier() string
	Load(context.Context, *config.Config) error
}

// BaseResource is the default implementation for Resource
type BaseResource struct {
}

// GetIdentifier returns a unique identifier for the resource.
func (base BaseResource) GetIdentifier() string {
	return "baseResource"
}

// Load is the method to load the resource.
func (base BaseResource) Load(context.Context, *config.Config) error {
	return nil
}
