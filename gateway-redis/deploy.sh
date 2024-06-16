#!/bin/bash

kubectl apply -f namespace.yaml
kubectl apply -f secrets.yaml
kubectl apply -f statefulset.yaml
kubectl apply -f service.yaml

# Set the namespace
NAMESPACE="gateway-redis"

# Get the IP addresses of the pods
POD_IPS_WITH_PORT=''
for i in {0..5}; do
    POD_NAME=redis-cluster-$i
    POD_IPS_WITH_PORT+=$(kubectl exec -n "$NAMESPACE" "$POD_NAME" -- ifconfig | grep 'inet addr:' | grep -v '127.0.0.1' | awk '{print $2}' | awk -F: '{print $2}')
    POD_IPS_WITH_PORT+=":6379 "
done

echo $POD_IPS_WITH_PORT

sleep 10

# Convert array to JSON array and create a ConfigMap
kubectl delete configmap redis-ips  -n "$NAMESPACE"
kubectl create configmap redis-ips --from-literal=ips="$(jq -r -n --arg ips "${POD_IPS_WITH_PORT[*]}" '$ips')" -n "$NAMESPACE"
kubectl get -o yaml configmap redis-ips  -n "$NAMESPACE"

# Display a message indicating the ConfigMap creation
echo "ConfigMap 'redis-ips' created with the following IP addresses:"

kubectl delete pods/redis-cluster-joiner -n "$NAMESPACE"
kubectl apply -f redis-cluster-joiner.yaml

# redis-cli --cluster create <IP_ADDR1>:6379 <IP_ADDR2>:6379 ... --cluster-replicas 1
