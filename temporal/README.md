# Temporal Setup
This directory contains the necessary configurations and scripts to set up Temporal in your Minikube environment.

# Directory Structure
```
.
├── README.md                  # This file
├── configmaps.yaml            # ConfigMaps required for Temporal
├── deploy.sh                  # Script to deploy Temporal
├── deployments.yaml           # Kubernetes deployments for Temporal services
├── namespace.yaml             # Namespace for Temporal
├── serviceaccount.yaml        # ServiceAccounts required for Temporal
└── services.yaml              # Kubernetes services for Temporal
```

# Setting Up Temporal
To set up Temporal in your Minikube environment, follow these steps:

## Clone the Repository:

```
git clone https://github.com/your-repo/temporal-setup.git
cd temporal-setup
```

## Deploy Temporal:
```
./deploy.sh
```

This script will:

1. Create the namespace.
2. Apply ConfigMaps.
3. Create ServiceAccounts.
4. Deploy Temporal services.
5. Expose the services.

## Accessing Temporal
Once deployed, you can access the Temporal frontend service using the service IP or DNS name provided by your Minikube environment. You can also port-forward to access it locally:

```
kubectl port-forward svc/temporal-frontend 7233:7233 -n temporal-namespace
```