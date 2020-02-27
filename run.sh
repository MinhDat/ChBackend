#!/bin/bash
export PORT=9000
export PG_USER=postgres
export PG_PASSWORD=c9BqhGZM5v7EPTs7
export PG_HOST=localhost
export PG_PORT=5432
export PG_DB=tasks
if [ ! "$(docker ps -aq -f status=running -f name=postgres)" ]; then
    docker-compose -f docker-compose-postgres.yml up -d
fi
go run main.go
