#!/bin/bash

# Script to force a specific database migration version
# Usage: ./fix-dirty-db-migrate.sh <version>

if [ -z "$1" ]; then
  echo "Usage: ./force-migrate.sh <version>"
  echo "Example: ./force-migrate.sh 11"
  exit 1
fi

VERSION=$1

# Load environment variables from .env (skip comments and empty lines, remove inline comments)
export $(grep -v '^#' .env | grep -v '^$' | sed 's/ #.*//' | xargs)

# Build connection string
DB_CONNECTION="postgres://${DB_USER}:${PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

# Force migrate to the specified version
migrate -path ./migrations -database "$DB_CONNECTION" force "$VERSION"
