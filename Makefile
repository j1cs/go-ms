build:
	go build -o cmd/api main.go

run:
	go run cmd/api/main.go

dl:
	go mod download