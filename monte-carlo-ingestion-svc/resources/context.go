package resources

import "context"

// Context is an abstraction to hold the top level context
type Context interface {
	GetContext() context.Context
	Cancel()
}

// ContextImpl implements Context
type ContextImpl struct {
	ctx    context.Context
	cancel context.CancelFunc
}

// NewContext is the constructor for context
func NewContext() *ContextImpl {
	ctx, cancel := context.WithCancel(context.Background())
	return &ContextImpl{ctx, cancel}
}

// GetContext returns the context
func (context *ContextImpl) GetContext() context.Context {
	return context.ctx
}

// Cancel cancels the context
func (context *ContextImpl) Cancel() {
	context.cancel()
}

// GetIdentifier returns the resource identifier
func (context *ContextImpl) GetIdentifier() string {
	return "context"
}