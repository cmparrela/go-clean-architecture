build-api: 
	go build -o ./bin/api api/main.go

build-cmd:
	go build -o ./bin/cmd cmd/main.go

run-api:
	go run api/main.go

run-cmd:
	go run cmd/main.go