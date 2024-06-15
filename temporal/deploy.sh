#!/bin/bash

kubectl apply -f namespace.yaml
kubectl apply -f configmaps.yaml
kubectl apply -f serviceaccount.yaml
kubectl apply -f deployments.yaml
kubectl apply -f services.yaml