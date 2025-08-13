#!/bin/bash
set -e

echo "Starting portfolio..."
cd portfolio && docker compose down && docker compose up --build -d
cd ..

echo "Starting nginx-proxy..."
cd nginx-proxy && docker compose down && docker compose up --build -d