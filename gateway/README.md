# Kong Gateway Setup
This directory contains the necessary configurations and scripts to set up the Kong Gateway with a Redis cluster and a custom authentication plugin in your Minikube environment.

# Directory Structure
```
.
├── Dockerfile                     # Dockerfile for Kong setup
├── README.md                      # This file
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
