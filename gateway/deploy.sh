#!/bin/bash

cd authentication-plugin/
luarocks make
luarocks pack kong-plugin-authentication '1.0-1'
cd ..

rm -rf plugins
mkdir plugins

mv authentication-plugin/kong-plugin-authentication-1.0-1.all.rock plugins/

eval $(minikube docker-env | grep -v "│" | grep -v "─")

docker build -t gateway .

kubectl apply -f configmaps.yaml
kubectl apply -f namespace.yaml
kubectl apply -f secrets.yaml
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
