#!/bin/bash

# Start the application and dependencies
docker-compose up -d --build --force-recreate

# Wait for the application to start
echo "Waiting for application to start"
until curl -s http://localhost:8080/healthcheck; do
    sleep 1
done
echo "Application started successfully"

go run ./test/integration/main.go


docker-compose down