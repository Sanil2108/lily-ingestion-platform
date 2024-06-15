#!/bin/bash
$(./scripts/generate-wire.sh)
go build && exec $(./monte-carlo-ingestion)
