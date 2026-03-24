.PHONY: run build clean

APP_DIR := services/application
BINARY := bin/server

run:
	cd $(APP_DIR) && go run ./cmd/server

test:
	cd $(APP_DIR) && go test ./..

build:
	cd $(APP_DIR) && go build -o ../../$(BINARY) ./cmd/server

clean:
	rm -f $(BINARY)
