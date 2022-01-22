#!make
include .env

all: build run

run:
	docker-compose up -d
stop:
	docker-compose down
build:
	docker-compose build

DB_URL := postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DB}?sslmode=disable
migrate:
	@migrate -database $(DB_URL) -path db/migrations up
migrate-down:
	@migrate -database $(DB_URL) -path db/migrations down
