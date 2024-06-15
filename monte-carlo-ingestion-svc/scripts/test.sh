#!/bin/bash
$(./scripts/generate-wire.sh)
go clean -testcache

go test \
  -v ./... \
  -coverprofile=coverage.out \
  -coverpkg ./resources/...,./common/...,./controllers/...
