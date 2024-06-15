package controllers

import (
	"context"

	"monte-carlo-ingestion/domain"
	"monte-carlo-ingestion/services"
)

type IngestionController struct {
	DefaultController
	ingestionService services.IngestionService
}

func (ingestionController *IngestionController) NewRequest() *domain.Request {
	return &domain.Request{
		Body: &domain.IngestionRequest{},
	}
}

func (ingestionController *IngestionController) Handler(ctx context.Context, req *domain.Request) (interface{}, error) {
	ingestionRequest := req.Body.(*domain.IngestionRequest)

	response, err := ingestionController.ingestionService.Ingest(*ingestionRequest)

	return response, err
}

func NewIngestionController(ingestionService services.IngestionService) IngestionController {
	return IngestionController{
		ingestionService: ingestionService,
	}
}
