#!/bin/bash

set -e

NETWORK_NAME="company_network"

if ! podman network exists $NETWORK_NAME; then
  podman network create $NETWORK_NAME
fi

if ! podman image exists service:latest; then
  podman build -t service .
fi

podman run -d --replace \
  --name cockroachdb \
  --network $NETWORK_NAME \
  -p 26257:26257 \
  -p 8080:8080 \
  -v cockroachdb-data:/cockroach/cockroach-data \
  -v ./migrations/001_init.sql:/docker-entrypoint-initdb.d/001_init.sql \
  cockroachdb/cockroach:v25.1.2 start-single-node --insecure

echo "Waiting for CockroachDB to become healthy..."
for i in {1..5}; do
  if podman exec cockroachdb curl -f http://localhost:8080/health?ready=1 >/dev/null 2>&1; then
    echo "CockroachDB is healthy!"
    break
  fi
  echo "Retrying... ($i/5)"
  sleep 10
done

podman run --replace \
  --name service \
  --network $NETWORK_NAME \
  -p 8081:8081 \
  --env DATABASE_HOST=cockroachdb \
  service:latest
