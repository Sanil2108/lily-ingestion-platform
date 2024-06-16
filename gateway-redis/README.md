
# Gateway Redis Cluster Setup
This directory contains the necessary configurations and scripts to set up a Redis cluster for the Gateway in your Minikube environment.

# Directory Structure
```
.
├── README.md                  # This file
├── deploy.sh                  # Script to deploy Redis cluster
├── namespace.yaml             # Namespace for Redis cluster
├── redis-cluster-joiner.yaml  # Pod to join Redis nodes into a cluster
├── secrets.yaml               # Secrets for Redis cluster
├── service.yaml               # Kubernetes service for Redis
└── statefulset.yaml           # StatefulSet for Redis nodes
```

# Setting Up Redis Cluster
To set up the Redis cluster in your Minikube environment, follow these steps:

## Clone the Repository:

```
git clone https://github.com/your-repo/gateway-redis-setup.git
cd gateway-redis-setup
```

## Deploy Redis Cluster:

```
./deploy.sh
```

This script will:

1. Create the namespace.
2. Apply secrets.
3. Create the StatefulSet for Redis nodes.
4. Create the service.
5. Get the IP addresses of the Redis pods.
6. Create a ConfigMap with these IP addresses.
7. Deploy a pod (redis-cluster-joiner) to join the Redis nodes into a cluster.

# How the Redis Cluster is Formed

Namespace Creation:
The namespace.yaml file defines the namespace for the Redis cluster.

Secrets Application:
The secrets.yaml file contains necessary secrets for the Redis cluster.

StatefulSet Deployment:
The statefulset.yaml file defines the StatefulSet for Redis nodes, ensuring that each pod gets a stable network identity and persistent storage.

Service Creation:
The service.yaml file creates a headless service to manage the Redis pods.

Fetching Pod IPs:
The deploy.sh script fetches the IP addresses of all Redis pods and creates a ConfigMap (redis-ips) containing these IP addresses.

Cluster Joining:
The redis-cluster-joiner.yaml file defines a pod that reads the IP addresses from the ConfigMap and runs the redis-cli --cluster create command to join the Redis nodes into a cluster.
