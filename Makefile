# Run application in development mode
dev:
	air
# Generate or update swagger docs
swag:
	swag init --dir ./cmd/server,./internal/handlers,./internal/services,./internal/routes,./internal/payments/proto