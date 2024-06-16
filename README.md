# Atlan Lily

This repository contains the implementation of a metadata ingestion system, designed to work with multiple integrations and orchestrate workflows using Temporal. The key components included in this project are:

[Architecture of the system (with the implemented components highlighted)](https://i.imgur.com/crDXufh.png)

1. Monte Carlo Ingestion Service: A service for ingesting health metadata from Monte Carlo.
2. Temporal Cluster Setup: Configuration files and scripts for deploying Temporal in a Minikube environment.
3. Kong Gateway: Setup and configuration for Kong Gateway, including a custom authentication plugin.
4. Redis Cluster: Setup and deployment scripts for a Redis cluster to be used with Kong Gateway.

## Monte Carlo Ingestion Service
Responsible for receiving health metadata updates from Monte Carlo via webhooks, processing them, and initiating workflows using Temporal for further data propagation.

## Temporal
Orchestration platform for executing workflows in a distributed manner. Handles task execution, retries, and workflow state management across various components.

## Gateway Redis
Manages Kong's API gateway configuration and plugins, ensuring scalable authentication and authorization functionalities within the microservices architecture.

## Kong Gateway
API gateway orchestrating access to services, integrating with Redis for efficient request routing, rate limiting, and security enforcement.

