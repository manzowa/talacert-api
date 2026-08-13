APP_NAME=talacert-api
MAIN_PATH=cmd/main.go

ifeq ($(OS),Windows_NT)
    EXE=.exe
    RM=if exist bin rmdir /s /q bin
else
    EXE=
    RM=rm -rf bin
endif

.PHONY: all build run clean test tidy swagger install fmt vet

all: run

install:
	go mod download

tidy:
	go mod tidy

build:
	go build -o bin/$(APP_NAME)$(EXE) $(MAIN_PATH)

run:
	go run $(MAIN_PATH)

test:
	go test ./...

swagger:
	swag init -g $(MAIN_PATH)

clean:
	$(RM)

fmt:
	go fmt ./...

vet:
	go vet ./...