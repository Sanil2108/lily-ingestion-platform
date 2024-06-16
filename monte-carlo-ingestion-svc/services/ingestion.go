package services

import (
	"monte-carlo-ingestion/domain"
	"monte-carlo-ingestion/logger"
	"monte-carlo-ingestion/resources"
)

type IngestionService struct {
	logger   logger.Logger
	temporal resources.Temporal
}

func NewIngestionService(logger *logger.Logger, temporal resources.Temporal) IngestionService {
	return IngestionService{
		logger:   *logger,
		temporal: temporal,
	}
}

func (ingestionService IngestionService) Ingest(req domain.IngestionRequest) (domain.IngestionResponse, error) {
	ingestionService.logger.Log.Info("Ingesting data")

	workflowId, err := ingestionService.temporal.StartWorkflow("monte-carlo-ingestion", "ingestion-tq-"+req.TenantId, map[string]interface{}{})
	if err != nil {
		return domain.IngestionResponse{}, err
	}

	ingestionService.logger.Log.Info("Ingestion completed")

	response := domain.IngestionResponse{
		Status:     "ok",
		WorkflowId: workflowId,
	}
	return response, nil
}
