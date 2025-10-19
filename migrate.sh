#!/bin/sh

COMMAND=${1:-up}

echo "Running migration: $COMMAND"

if [ -f "./migrate" ]; then
    ./migrate migrate:$COMMAND
else
    echo "Migrate binary not found!"
    echo "Running migration using golang-migrate CLI..."

    if ! command -v migrate &> /dev/null; then
        echo "Installing golang-migrate..."
        apk add --no-cache curl
        curl -L https://github.com/golang-migrate/migrate/releases/download/v4.19.0/migrate.linux-amd64.tar.gz | tar xvz
        mv migrate /usr/local/bin/migrate
        chmod +x /usr/local/bin/migrate
    fi

    migrate -path=./database/migrations -database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" $COMMAND
fi