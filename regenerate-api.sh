#!/bin/bash

set -e  # Exit on error

echo "Regenerating API documentation..."
cd backend
swag init


echo ""
echo "Updating frontend API client..."
cd ../frontend
echo "Running orval to generate API client..."
npx orval
