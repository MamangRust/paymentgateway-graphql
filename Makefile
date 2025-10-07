migrate:
	go run cmd/migrate/main.go up

migrate-down:
	go run cmd/migrate/main.go down

graphql-generate:
	gqlgen generate

sqlc-generate:
	sqlc generate

run-server:
	go run cmd/server/main.go