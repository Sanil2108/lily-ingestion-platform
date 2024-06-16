package resources

import (
	"context"

	"monte-carlo-ingestion/config"
	"monte-carlo-ingestion/logger"

	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
)

// BaseResource is the default implementation for Resource
type Temporal interface {
	Resource

	StartWorkflow(workflowName string, taskQueue string, input map[string]interface{}) (string, error)
}

func NewTemporal(logger *logger.Logger) (*TemporalImpl, error) {
	client, err := client.NewClient(client.Options{})
	if err != nil {
		return nil, err
	}
	return &TemporalImpl{
		client: client,
		logger: *logger,
	}, nil
}

type TemporalImpl struct {
	BaseResource

	client client.Client
	logger logger.Logger
}

// GetIdentifier returns a unique identifier for the resource.
func (temporal TemporalImpl) GetIdentifier() string {
	return "temporal"
}

// Load is the method to load the resource.
func (temporal TemporalImpl) Load(context.Context, *config.Config) error {
	return nil
}

func (temporal TemporalImpl) StartWorkflow(
	workflowName string,
	taskQueue string,
	input map[string]interface{},
) (string, error) {
	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: taskQueue,
	}

	temporal.logger.Log.Info("Starting workflow", zap.String("workflowName", workflowName))

	workflowRun, err := temporal.client.ExecuteWorkflow(context.Background(), workflowOptions, workflowName)
	if err != nil {
		return "", err
	}
	temporal.logger.Log.Info(
		"Starting workflow",
		zap.String("workflowId", workflowRun.GetID()),
		zap.String("workflowRunId", workflowRun.GetRunID()),
	)

	return workflowRun.GetID(), nil
}
