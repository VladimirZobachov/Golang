
	run:
		@go run cmd/hostess-service/main.go

build:
	@go build cmd/hostess-service/main.go

swag:
	@swag init -g cmd/hostess-service/main.go

.PHONY: run build swag

