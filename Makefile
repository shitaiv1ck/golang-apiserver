include .env
export

app-run:
	@go run cmd/main.go

app-build:
	@go build -v ./cmd 

migrate-up:
	@migrate -path ./migrations -database \
	"postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" \
	up

migrate-down:
	@migrate -path ./migrations -database \
	"postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" \
	down

migrate-force-one:
	@migrate -path ./migrations -database \
	"postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" \
	force 1