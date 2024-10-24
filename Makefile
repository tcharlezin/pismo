APP_BINARY=pismo

up_app: build_app
	docker-compose --env-file .env  up -d

## up: stops docker-compose (if running), builds all projects and starts docker compose
up: build_app
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose --env-file .env  up --build -d
	@echo "Docker images built and started!"

## build_app: builds the app binary as a linux executable
build_app:
	@echo "Building broker binary..."
	env GOOS=linux CGO_ENABLED=0 go build -o ${APP_BINARY} ./cmd/
	@echo "Done!"

## clear: remove the app binary
clear:
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Removing binary..."
	rm ${APP_BINARY}
	@echo "Done!"

## doc: generate documentation
doc:
	@echo "Generating docs..."
	@swag init -d "./" -g "cmd/main.go"
	@echo "...Done!"

## migrate: run migrations
migrate:
	@echo "Setup migrations..."
	@migrate -path=internal/database/migrations -database "postgresql://postgres:password@localhost:5432/pismo?sslmode=disable" -verbose up
	@echo "...Done!"


## rollback: rollback migrations
rollback:
	@echo "Rollback migrations..."
	@migrate -path=internal/database/migrations -database "postgresql://postgres:password@localhost:5432/pismo?sslmode=disable" -verbose down
	@echo "...Done!"

## test-unit: test unit
test-unit:
	@echo "Test unit..."
	go test ./...
	@echo "...Done!"