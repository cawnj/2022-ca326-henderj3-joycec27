#!make
include .env

all: test build run

run:
	docker-compose up -d
stop:
	docker-compose down
build:
	docker-compose build
test:
	go test -count 1 ./...

DB_URL := postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DB}?sslmode=disable
migrate:
	@docker run \
	-v `pwd`/db/migrations:/migrations \
	--network host migrate/migrate \
	-database $(DB_URL) \
	-path /migrations \
	up
migrate-down:
	@docker run \
	-v `pwd`/db/migrations:/migrations \
	--network host migrate/migrate \
	-database $(DB_URL) \
	-path /migrations \
	down -all
