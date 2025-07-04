.PHONY: migrate migrate-down swagger

migrate:
	go run cmd/migrations/main.go -up

migrate-down:
	go run cmd/migrations/main.go -down

# Generate Swagger documentation
swagger:
	$(HOME)/go/bin/swag init -g cmd/main.go

# Run the application
run:
	go run cmd/main.go

# Generate Swagger docs and run the application
dev: swagger run