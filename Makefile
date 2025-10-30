run:
	@go run cmd/server/main.go

docs:
	@swag init -g cmd/server/main.go