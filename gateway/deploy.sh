#!/bin/bash

eval $(minikube docker-env | grep -v "│" | grep -v "─")

docker build -t gateway .

kubectl apply -f configmaps.yaml
kubectl apply -f namespace.yaml
kubectl apply -f secrets.yaml
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
