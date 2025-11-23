.PHONY: help dev dev-down prod prod-down backend frontend clean test lint docker-build docker-push

# Default target
help:
	@echo "Crypto Marketplace - Available Commands:"
	@echo ""
	@echo "Development:"
	@echo "  make dev              - Start development environment with hot-reloading"
	@echo "  make dev-down         - Stop development environment"
	@echo "  make backend          - Run backend only (local)"
	@echo "  make frontend         - Run frontend only (local)"
	@echo ""
	@echo "Production:"
	@echo "  make prod             - Start production environment with Docker"
	@echo "  make prod-down        - Stop production environment"
	@echo ""
	@echo "Testing & Quality:"
	@echo "  make test             - Run all tests"
	@echo "  make test-backend     - Run backend tests"
	@echo "  make test-frontend    - Run frontend tests"
	@echo "  make lint             - Run linters"
	@echo ""
	@echo "Docker:"
	@echo "  make docker-build     - Build Docker images"
	@echo "  make docker-push      - Push Docker images to registry"
	@echo ""
	@echo "Utilities:"
	@echo "  make clean            - Clean build artifacts and caches"
	@echo "  make logs             - Show logs from all services"
	@echo "  make health           - Check health of services"

# Development
dev:
	@echo "Starting development environment..."
	docker-compose -f docker-compose.dev.yml up --build

dev-down:
	@echo "Stopping development environment..."
	docker-compose -f docker-compose.dev.yml down

backend:
	@echo "Starting backend..."
	cd backend && go run main.go

frontend:
	@echo "Starting frontend..."
	cd frontend && npm run dev

# Production
prod:
	@echo "Starting production environment..."
	docker-compose up --build -d
	@echo "Services started. Access:"
	@echo "  Frontend: http://localhost:3000"
	@echo "  Backend API: http://localhost:8080/api"
	@echo "  Backend Health: http://localhost:8080/health"

prod-down:
	@echo "Stopping production environment..."
	docker-compose down

# Testing
test: test-backend test-frontend

test-backend:
	@echo "Running backend tests..."
	cd backend && go test -v -race -coverprofile=coverage.out ./...

test-frontend:
	@echo "Running frontend tests..."
	cd frontend && npm test

# Linting
lint:
	@echo "Running linters..."
	@echo "Backend: go vet & go fmt"
	cd backend && go vet ./...
	cd backend && gofmt -s -l .
	@echo "Frontend: npm run lint"
	cd frontend && npm run lint --if-present

# Docker
docker-build:
	@echo "Building Docker images..."
	docker build -t crypto-marketplace-backend:latest ./backend
	docker build -t crypto-marketplace-frontend:latest ./frontend

docker-push:
	@echo "Pushing Docker images..."
	docker push crypto-marketplace-backend:latest
	docker push crypto-marketplace-frontend:latest

# Utilities
clean:
	@echo "Cleaning build artifacts..."
	cd backend && rm -f crypto-marketplace coverage.out coverage.html
	cd backend && rm -rf tmp
	cd frontend && rm -rf .nuxt .output dist node_modules/.cache
	docker-compose down -v --remove-orphans
	docker-compose -f docker-compose.dev.yml down -v --remove-orphans

logs:
	docker-compose logs -f

health:
	@echo "Checking backend health..."
	@curl -f http://localhost:8080/health || echo "Backend is not healthy"
	@echo ""
	@echo "Checking frontend health..."
	@curl -f http://localhost:3000 || echo "Frontend is not healthy"

# Install dependencies
install:
	@echo "Installing backend dependencies..."
	cd backend && go mod download
	@echo "Installing frontend dependencies..."
	cd frontend && npm install

# Database migrations (placeholder for future)
migrate-up:
	@echo "Running migrations..."
	# Add migration commands here

migrate-down:
	@echo "Rolling back migrations..."
	# Add rollback commands here
