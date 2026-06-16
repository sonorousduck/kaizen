#!/bin/bash
# Wrapper around `migrate create` from golang-migrate.

MIGRATIONS_DIR="migrations"
mkdir -p "$MIGRATIONS_DIR"

read -rp "Migration name (e.g. create_users_table): " name

if [[ -z "$name" ]]; then
    echo "Error: migration name cannot be empty"
    exit 1
fi

migrate create -ext sql -dir "$MIGRATIONS_DIR" -seq "$name"
