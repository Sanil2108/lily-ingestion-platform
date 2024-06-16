package mockresources

import (
	"monte-carlo-ingestion/resources"
)

type TemporalImpl struct {
	resources.BaseResource
}

func (temporal TemporalImpl) StartWorkflow(
	workflowName string,
	taskQueue string,
	input map[string]interface{},
) (string, error) {
	return "", nil
}

func NewTemporal() (*TemporalImpl, error) {
	return &TemporalImpl{}, nil
}
