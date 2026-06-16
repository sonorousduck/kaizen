#!/bin/bash

# Regenerate OpenAPI spec from Go comments and update frontend API client
# Run this from the velocity directory

set -e  # Exit on error

echo "Regenerating API documentation..."
cd backend
swag init

cd docs

# Check if swagger2openapi is installed, install if needed
if ! command -v swagger2openapi &> /dev/null; then
    echo "Installing swagger2openapi..."
    pnpm install -g swagger2openapi > /dev/null 2>&1
fi

# Convert swagger.json to openapi.json (OpenAPI 3.0 format)
swagger2openapi swagger.json > openapi.json

echo " Backend OpenAPI spec generated"

echo ""
echo "Updating frontend API client..."
cd ../../frontend
echo "Running orval to generate API client..."
npx orval
echo "Frontend API client generated"

echo "Done! API has been regenerated successfully"
