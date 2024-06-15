package mockresources

import (
	"monte-carlo-ingestion/resources"
)

type TemporalImpl struct {
	resources.BaseResource
}

func (temporal TemporalImpl) StartWorkflow(
	workflowName string,
	input map[string]interface{},
) error {
	return nil
}

func NewTemporal() (*TemporalImpl, error) {
	return &TemporalImpl{}, nil
}
