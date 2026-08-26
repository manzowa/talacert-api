APP_NAME := talacert-api
MAIN_PATH := cmd/main.go
BIN_DIR := bin

ifeq ($(OS),Windows_NT)
    EXE := .exe
    RM := if exist $(BIN_DIR) rmdir /s /q $(BIN_DIR)
    START := $(BIN_DIR)/$(APP_NAME)$(EXE)
else
    EXE :=
    RM := rm -rf $(BIN_DIR)
    START := ./$(BIN_DIR)/$(APP_NAME)
endif

.PHONY: all build start run clean clean-go-cache clean-all test tidy swagger install fmt vet

all: start

install:
	go mod download

tidy:
	go mod tidy

build:
	go build -o $(BIN_DIR)/$(APP_NAME)$(EXE) $(MAIN_PATH)

start: build
	$(START)

run:
	go run $(MAIN_PATH)

test:
	go test ./...

swagger:
	swag init -g $(MAIN_PATH) --parseInternal --parseDependency

clean:
	$(RM)

clean-go-cache:
	go clean -cache
	go clean -modcache

clean-all: clean clean-go-cache

fmt:
	go fmt ./...

vet:
	go vet ./...
