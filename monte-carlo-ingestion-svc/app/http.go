// Package app provides the implementation of a self-contained HTTP server.
// It includes various controllers for handling different types of requests.
// The server can be configured using a config file and uses resources like SQS.
// It also uses a logger for logging purposes.
package app

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"monte-carlo-ingestion/config"
	"monte-carlo-ingestion/controllers"
	"monte-carlo-ingestion/logger"
	"monte-carlo-ingestion/resources"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Controllers holds all request handlers
type Controllers struct {
	health     controllers.HealthController
	ingestion controllers.IngestionController
}

// Resources holds all the resources required by the application
type Resources struct {
	temporal resources.Temporal
}

// HTTPApp is the self contained http server
type HTTPApp struct {
	cfg         *config.Config
	logger      *logger.Logger
	boot        Bootstrap
	resources   Resources
	controllers Controllers
}

// NewHTTPApp helps set up a new instance of HTTPApp.
func NewHTTPApp(
	config *config.Config,

	ctx resources.Context,

	boot Bootstrap,
	health controllers.HealthController,
	ingestion controllers.IngestionController,
	logger *logger.Logger,
	temporal resources.Temporal,
) App {
	return &HTTPApp{
		cfg:  config,
		boot: boot,
		resources: Resources{
			temporal: temporal,
		},
		controllers: Controllers{
			health:     health,
			ingestion: ingestion,
		},
		logger: logger,
	}
}

// Load loads the necessary resources for the HTTPApp.
func (app *HTTPApp) Load() {
	errors := app.boot.Load(
		WithResource(app.resources.temporal),
	)
	if len(errors) > 0 {
		logger.Fatal("Failed to load", zap.Any("errors", errors))
	}
}

// Start runs the HTTP server.
func (app *HTTPApp) Start() {
	router := gin.New()
	router.Use(gin.Recovery())

	meta := router.Group("/_meta")
	{
		meta.GET("/health", controllers.HTTPController(&app.controllers.health))
	}

	api := router.Group("/_api")
	{
		api.POST("/ingestion", controllers.HTTPController(&app.controllers.ingestion))
	}

	logger.Info(fmt.Sprintf("Starting the application on pid %d in %s mode", os.Getpid(), app.cfg.Env))

	srv := &http.Server{
		Addr:        fmt.Sprintf("%s:%d", app.cfg.Server.Host, app.cfg.Server.Port),
		Handler:     router,
		IdleTimeout: 90 * time.Second,
	}
	srv.ListenAndServe()
	logger.Info("Server exiting")
}
