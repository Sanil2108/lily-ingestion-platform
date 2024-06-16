# Kong Gateway Setup
This directory contains the necessary configurations and scripts to set up the Kong Gateway with a Redis cluster and a custom authentication plugin in your Minikube environment.

# Directory Structure
```
.
├── Dockerfile                     # Dockerfile for Kong setup
├── README.md                      # This file
├── authentication-plugin          # Custom authentication plugin for Kong
│   ├── handler.lua                # Plugin handler code
│   ├── kong-plugin-authentication-1.0-1.rockspec # Plugin specification
│   └── schema.lua                 # Plugin schema definition
├── configmaps.yaml                # ConfigMaps for Kong configuration
├── deploy.sh                      # Script to deploy Kong Gateway and plugins
├── deployment.yaml                # Deployment configuration for Kong
├── namespace.yaml                 # Namespace for Kong
├── plugins                        # Directory for Kong plugin package
│   └── kong-plugin-authentication-1.0-1.all.rock # Compiled plugin
├── secrets.yaml                   # Secrets for Kong configuration
└── service.yaml                   # Service configuration for Kong
```

# Setting Up Kong Gateway
To set up the Kong Gateway in your Minikube environment, follow these steps:

Clone the Repository:

```
git clone https://github.com/your-repo/kong-gateway-setup.git
cd kong-gateway-setup
```

# Deploy Kong Gateway:

```
./deploy.sh
```

This script will:

1. Build the custom authentication plugin using Luarocks.
2. Package the custom authentication plugin and move it to the plugins directory.
3. Set up the Docker environment for Minikube.
4. Build the Docker image for the Kong Gateway.
5. Apply the ConfigMaps for Kong configuration.
6. Create the namespace for Kong.
7. Apply the secrets for Kong configuration.
8. Deploy the Kong Gateway using the deployment.yaml file.
9. Create the service for Kong Gateway.