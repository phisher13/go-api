build:
	docker-compose build go-api

run:
	docker-compose up go-api

migrate:
	migrate -path ./migrate -database 'postgres://postgres:root@0.0.0.0:5433/postgres?sslmode=disable' -verbose up

.PHONY: build run migrate