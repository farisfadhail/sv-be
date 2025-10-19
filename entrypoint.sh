#!/bin/sh

echo "Starting application..."

echo "Waiting for MySQL to be ready..."
sleep 5

echo "Running database migrations..."
./migrate migrate:up

if [ "$RUN_SEEDER" = "true" ]; then
    echo "Running database seeders..."
    ./seeder
fi

echo "Starting main application..."
exec ./main