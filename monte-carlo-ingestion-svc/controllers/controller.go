package controllers

import (
	"context"

	"monte-carlo-ingestion/domain"

	customError "monte-carlo-ingestion/errors"
	"monte-carlo-ingestion/logger"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// Controller Interface represents the contract that a controller must adhere to.
type Controller interface {
	// NewRequest creates a new instance of the request object.
	NewRequest() *domain.Request

	// ValidateRequest validates the incoming request.
	ValidateRequest(context.Context, *domain.Request) error

	// Handler handles the incoming request and returns the response.
	Handler(context.Context, *domain.Request) (interface{}, error)
}


type DefaultController struct {
}

var validate = validator.New()

// NewRequest is the constructor for the request.
func (c *DefaultController) NewRequest() *domain.Request {
	return &domain.Request{}
}

// ValidateRequest is the default implementation to validate the incoming request.
func (c *DefaultController) ValidateRequest(_ context.Context, req *domain.Request) error {
	if req.Body != nil {
		if err := validate.Struct(req.Body); err != nil {
			return customError.NewBadRequestError(customError.WithMsg("request body " + err.Error()))
		}
	}

	if req.QueryParams != nil {
		if err := validate.Struct(req.QueryParams); err != nil {
			logger.Error("Error validating query", zap.Error(err))
			return customError.NewBadRequestError(customError.WithError("request query " + err.Error()))
		}
	}

	return nil
}

// Handler is the default implementation for the handler of the incoming request.
func (c *DefaultController) Handler(_ context.Context, _ *domain.Request) (*domain.Response, error) {
	return nil, nil
}
