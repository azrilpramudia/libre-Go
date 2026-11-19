-include .env
export

.PHONY: run migrate-up migrate-down migrate-create sqlc-generate docker-up docker-down

run:
	air

migrate-up:
	migrate -path db/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" up

migrate-down:
	migrate -path db/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" down 1

migrate-create:
	migrate create -ext sql -dir db/migrations -seq $(name)

sqlc-generate:
	sqlc generate

docker-up:
	docker compose --env-file .env -f ../docker-compose.yml -p backend up -d --build

docker-down:
	docker compose --env-file .env -f ../docker-compose.yml -p backend down
