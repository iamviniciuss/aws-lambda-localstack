#!/bin/bash
GOOS=linux GOARCH=amd64 go build -o handler src/handler.go
serverless deploy --stage local --verbose