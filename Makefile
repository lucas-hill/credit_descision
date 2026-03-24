.PHONY: infra-up infra-down db-setup dev run test build clean

APP_DIR := services/application
BINARY := bin/server

infra-up:
	docker compose up -d

infra-down:
	docker compose down

db-setup:
	./scripts/setup-dynamodb.sh

dev: infra-up db-setup run-app

run:
	cd $(APP_DIR) && go run ./cmd/server

test:
	cd $(APP_DIR) && go test ./..

build:
	cd $(APP_DIR) && go build -o ../../$(BINARY) ./cmd/server

clean:
	rm -f $(BINARY)
