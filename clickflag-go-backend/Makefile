# Makefile for ClickFlag Backend

.PHONY: help build run test clean docker-build docker-run docker-stop docker-logs logs-check logs-clean logs-status

# Default target
help:
	@echo "Available commands:"
	@echo "  build        - Build the application"
	@echo "  run          - Run the application locally"
	@echo "  test         - Run tests"
	@echo "  clean        - Clean build artifacts"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run Docker container"
	@echo "  docker-stop  - Stop Docker containers"
	@echo "  docker-logs  - Show Docker logs"
	@echo "  logs-check   - Check log directory status"
	@echo "  logs-clean   - Clean old log files"
	@echo "  logs-status  - Show log system status"

# Build the application
build:
	go build -o bin/server ./cmd/server

# Build for Linux (for VPS deployment)
build-linux:
	GOOS=linux GOARCH=amd64 go build -o main ./cmd/server

# Run the application locally
run:
	go run ./cmd/server/main.go

# Run tests
test:
	go test ./tests/...

# Clean build artifacts
clean:
	rm -rf bin/
	rm -rf tmp/

# Docker commands
docker-build:
	docker build -t clickflag-backend .

docker-run:
	docker-compose up -d

docker-stop:
	docker-compose down

docker-logs:
	docker-compose logs -f

# Environment setup
env-setup:
	@if [ ! -f .env ]; then cp env.example .env; echo "Created .env from env.example"; fi

# Log management commands
logs-check:
	@echo "=== Log Directory Status ==="
	@if [ -d "./logs" ]; then \
		echo "Directory: ./logs"; \
		echo "Total size: $$(du -sm ./logs 2>/dev/null | cut -f1 || echo '0')MB"; \
		echo "Log files: $$(find ./logs -name '*.log*' 2>/dev/null | wc -l)"; \
		echo ""; \
		echo "=== Log Files ==="; \
		ls -lah ./logs/*.log* 2>/dev/null || echo "No log files found"; \
	else \
		echo "Log directory does not exist!"; \
	fi

logs-clean:
	@echo "=== Cleaning Old Logs ==="
	@if [ -d "./logs" ]; then \
		echo "Removing log files older than 7 days..."; \
		find ./logs -name "*.log.*" -mtime +7 -delete 2>/dev/null || true; \
		echo "Keeping only the 10 most recent log files..."; \
		ls -t ./logs/*.log.* 2>/dev/null | tail -n +11 | xargs -r rm 2>/dev/null || true; \
		echo "Cleanup completed!"; \
	else \
		echo "Log directory does not exist!"; \
	fi

logs-status:
	@echo "=== Log System Status ==="
	@if pgrep -f "clickflag-backend" > /dev/null 2>&1; then \
		echo "✅ Application is running"; \
	else \
		echo "❌ Application is not running"; \
	fi
	@if [ -d "./logs" ]; then \
		echo "✅ Log directory exists"; \
		echo "📊 Log directory size: $$(du -sm ./logs 2>/dev/null | cut -f1 || echo '0')MB"; \
	else \
		echo "❌ Log directory does not exist"; \
	fi
	@if docker ps | grep -q "clickflag-backend" 2>/dev/null; then \
		echo "✅ Docker container is running"; \
		echo "📊 Docker logs size: $$(docker logs clickflag-backend 2>/dev/null | wc -c | awk '{print int($$1/1024/1024)}')MB"; \
	else \
		echo "❌ Docker container is not running"; \
	fi
