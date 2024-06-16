# Monte Carlo Ingestion Service

## Overview

The Monte Carlo Ingestion Service is responsible for ingesting health metadata from Monte Carlo into the Atlan Lily platform. This service processes health-related metadata updates, validates them, and triggers workflows to store and manage this data effectively. It integrates with the Temporal workflow orchestration system to handle various pre-ingestion, ingestion, and post-ingestion tasks.

## Directory Structure
```
.
├── README.md
├── app
│   ├── app.go          # Main application entry point
│   ├── boot.go         # Bootstrapping and initialization
│   └── http.go         # HTTP server setup
├── common
│   └── utils.go        # Utility functions
├── config
│   └── config.go       # Configuration management
├── controllers
│   ├── controller.go   # Base controller
│   ├── health.go       # Health check controller
│   ├── http.go         # HTTP controller setup
│   └── ingestion.go    # Ingestion-specific controller
├── coverage.out
├── domain
│   ├── request.go      # Request domain models
│   └── response.go     # Response domain models
├── errors
│   ├── base.go         # Base error types
│   ├── errors.go       # Custom error definitions
│   └── handlers.go     # Error handling middleware
├── go.mod
├── go.sum
├── logger
│   └── logger.go       # Logging setup and utilities
├── main.go             # Main file to run the application
├── mocks
│   └── temporal.go     # Mocks for testing
├── monte-carlo-ingestion
├── resources
│   ├── context.go      # Context management
│   ├── resource.go     # Resource definitions
│   └── temporal.go     # Temporal client and workflow management
├── scripts
│   ├── generate-wire.sh # Script for dependency injection code generation
│   ├── local.sh        # Script to run the service locally
│   └── test.sh         # Script to run tests
├── services
│   ├── health.go       # Health check service
│   └── ingestion.go    # Ingestion service logic
├── tests
│   └── controllers
│       ├── health_test.go  # Health check controller tests
│       ├── ingestion_test.go # Ingestion controller tests
│       └── utils.go     # Test utilities
└── wire
    ├── wire.go         # Dependency injection setup
    └── wire_gen.go     # Generated dependency injection code
```

## Responsibilities
1. Ingest Health Metadata: Ingest health-related metadata updates from Monte Carlo.
Process and Validate: Validate incoming metadata and ensure it meets the required format and standards.
2. Trigger Workflows: Use Temporal to manage workflows for handling metadata ingestion, including any pre- and post-processing tasks.
3. Error Handling: Manage errors effectively and provide meaningful feedback to clients.
4. Health Checks: Provide endpoints for checking the health and status of the service.

## Running Locally
To run the Monte Carlo Ingestion Service locally, follow these steps:

### Prerequisites
1. Go 1.16+ installed

### Steps

Clone the Repository:

```
git clone https://github.com/your-repo/monte-carlo-ingestion-service.git
cd monte-carlo-ingestion-service
```

### Install Dependencies:
```
go mod tidy
```

### Run locally
```
./scripts/local.sh
```

### To run the tests:

```
./scripts/test.sh
```

### Health Check
You can check the health of the service by accessing the health endpoint:

```
curl http://localhost:8080/health
```
