//go:build wireinject
// +build wireinject

package wire

import (
	"monte-carlo-ingestion/app"
	"monte-carlo-ingestion/config"
	"monte-carlo-ingestion/controllers"
	"monte-carlo-ingestion/logger"
	mockresources "monte-carlo-ingestion/mocks"
	"monte-carlo-ingestion/resources"
	"monte-carlo-ingestion/services"

	googlewire "github.com/google/wire"
)

func InitializeApplication() (app.App, error) {
	googlewire.Build(
		resources.NewContext,
		googlewire.Bind(new(resources.Context), new(*resources.ContextImpl)),
		config.NewConfig,
		logger.NewLogger,

		// ---------------------- Resources ----------------------
		resources.NewTemporal,
		googlewire.Bind(new(resources.Temporal), new(*resources.TemporalImpl)),

		// ---------------------- Controllers ----------------------
		controllers.NewHealthController,
		controllers.NewIngestionController,

		// ---------------------- Services ----------------------
		services.NewIngestionService,

		// ---------------------- App ----------------------
		app.NewBootstrap,
		app.NewHTTPApp,
	)
	return &app.HTTPApp{}, nil
}

func InitializaIngestionControllerWithMocks() (controllers.IngestionController, error) {
	googlewire.Build(
		logger.NewLogger,

		// ---------------------- Resources ----------------------
		mockresources.NewTemporal,
		googlewire.Bind(new(resources.Temporal), new(*mockresources.TemporalImpl)),

		// ---------------------- Controllers ----------------------
		controllers.NewIngestionController,

		// ---------------------- Services ----------------------
		services.NewIngestionService,
	)
	return controllers.IngestionController{}, nil
}

func InitializaHealthControllerWithMocks() (controllers.HealthController, error) {
	googlewire.Build(
		// ---------------------- Controllers ----------------------
		controllers.NewHealthController,
	)
	return controllers.HealthController{}, nil
}
