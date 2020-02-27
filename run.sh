#!/bin/bash
export PORT=9000
export PG_USER=postgres
export PG_PASSWORD=c9BqhGZM5v7EPTs7
export PG_HOST=localhost
export PG_PORT=5432
export PG_DB=tasks
# docker-compose -f docker-compose-postgres.yml
go run main.go
