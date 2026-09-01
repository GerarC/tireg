#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."

echo "Stopping stack and removing the Postgres volume..."
docker compose down -v

echo "Starting Postgres..."
docker compose up -d postgres

echo "Waiting for Postgres to be healthy..."
until [ "$(docker compose ps -q postgres | xargs docker inspect -f '{{.State.Health.Status}}')" = "healthy" ]; do
  sleep 1
done

echo "Postgres is up with a clean schema. GORM's AutoMigrate will recreate tables on the next app start."
