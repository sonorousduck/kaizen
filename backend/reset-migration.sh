#!/bin/bash

# Script to reset migrations: runs down then up
# Reads database configuration from .env file

set -e

# Get the directory where this script is located
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Load environment variables from .env
if [ ! -f "$SCRIPT_DIR/.env" ]; then
    echo "Error: .env file not found in $SCRIPT_DIR"
    exit 1
fi

# Parse the .env file manually to handle special characters
while IFS='=' read -r key value; do
    # Skip comments and empty lines
    [[ "$key" =~ ^#.*$ ]] && continue
    [[ -z "$key" ]] && continue
    # Remove inline comments and trim whitespace
    value="${value%% #*}"
    value="${value## }"
    value="${value%% }"
    export "$key=$value"
done < "$SCRIPT_DIR/.env"

# Build the postgres connection URL for golang-migrate
POSTGRES_URL="postgres://${DB_USER}:${PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=${SSL_MODE}"

echo "Resetting migrations..."
echo "Database: ${DB_NAME}"
echo ""

# Check if migrate is installed
if ! command -v migrate &> /dev/null; then
    echo "Error: 'migrate' command not found. Install it with:"
    echo "  go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"
    exit 1
fi

# Run down migrations
echo "Running down migrations..."
migrate -path "$SCRIPT_DIR/migrations" -database "$POSTGRES_URL" down -all 2>&1 || {
    if [ $? -eq 1 ]; then
        echo "Down migrations completed"
    else
        exit 1
    fi
}

echo ""
echo "Running up migrations..."
# Run up migrations
migrate -path "$SCRIPT_DIR/migrations" -database "$POSTGRES_URL" up

echo ""
echo "✓ Migrations reset successfully"
