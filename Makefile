GOPATH ?= $(HOME)/go

run-api:
	go run api/main.go

run-cmd:
	go run cmd/main.go

format:
	go fmt ./...

dev:
	DC_APP_ENV=dev $(GOPATH)/bin/reflex -s -r '\.go$$' make format run-api

test-cov:
	go test -coverprofile=cover.out ./... && go tool cover -html=cover.out -o cover.html