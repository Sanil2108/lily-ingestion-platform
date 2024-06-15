package app

import (
	"sync"

	"monte-carlo-ingestion/config"
	"monte-carlo-ingestion/logger"
	"monte-carlo-ingestion/resources"

	"go.uber.org/zap"
)

// Bootstrap helps load all resource and initialize the application
type Bootstrap struct {
	Context   resources.Context
	Config    *config.Config
	Resources []resources.Resource
	logger    *logger.Logger
}

// NewBootstrap is contructor for the Bootstrap
func NewBootstrap(context resources.Context, logger *logger.Logger, config *config.Config) Bootstrap {
	return Bootstrap{Config: config, Context: context, logger: logger}
}

// WithResource adds a resource to load as part of initial setup. All resources are loaded
// concurrently.
func WithResource(resource resources.Resource) func(*Bootstrap) {
	return func(boot *Bootstrap) {
		boot.Resources = append(boot.Resources, resource)
	}
}

// Load helps load all resources
func (boot *Bootstrap) Load(modifiers ...func(*Bootstrap)) []error {
	// Loading all the resources asynchronously
	var wg sync.WaitGroup
	wg.Add(len(boot.Resources))
	var errs []error
	var mutex sync.Mutex
	for _, rs := range boot.Resources {
		resource := rs
		go func() {
			defer wg.Done()
			if err := resource.Load(boot.Context.GetContext(), boot.Config); err != nil {
				mutex.Lock()
				defer mutex.Unlock()
				errs = append(errs, err)
			} else {
				logger.Info("✔ Successfully loaded resource", zap.String("resourceIdentifier", resource.GetIdentifier()))
			}
		}()
	}

	wg.Wait()
	return errs
}
