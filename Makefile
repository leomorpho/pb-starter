.PHONY: help dev dev-backend dev-frontend storybook build test lint format clean install deps check

# Default target
help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

# Development commands
dev: ## Start both backend and frontend in development mode
	@echo "Starting development environment..."
	@make -j2 dev-backend dev-frontend

dev-backend: ## Start PocketBase backend in watch mode
	@echo "Starting PocketBase backend..."
	@cd pb && modd

dev-frontend: ## Start SvelteKit frontend in development mode
	@echo "Starting SvelteKit frontend..."
	@cd sk && npm run dev

storybook: ## Start Storybook development server
	@echo "Starting Storybook..."
	@cd sk && npm run storybook

# Build commands
build: ## Build the entire project
	@echo "Building frontend..."
	@cd sk && npm run build
	@echo "Building backend..."
	@cd pb && go build
	@echo "Build complete!"

build-frontend: ## Build only the frontend
	@cd sk && npm run build

build-backend: ## Build only the backend
	@cd pb && go build

# Testing commands
test: ## Run all tests
	@cd sk && npm run test

test-unit: ## Run unit tests only
	@cd sk && npm run test:unit

test-e2e: ## Run e2e tests only
	@cd sk && npm run test:e2e

# Code quality commands
lint: ## Run linting
	@cd sk && npm run lint

format: ## Format code
	@cd sk && npm run format

check: ## Run type checking
	@cd sk && npm run check

# Setup commands
install: deps ## Install all dependencies

deps: ## Install dependencies for both frontend and backend
	@echo "Installing frontend dependencies..."
	@cd sk && npm install
	@echo "Installing backend dependencies..."
	@cd pb && go mod tidy
	@echo "Dependencies installed!"

# Production commands
preview: build-frontend ## Preview production build
	@cd sk && npm run preview

serve: build-frontend ## Serve the app with PocketBase
	@cd pb && ./pocketbase serve --http 0.0.0.0:8090

# Cleanup commands
clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	@rm -rf sk/build sk/.svelte-kit sk/node_modules/.vite
	@rm -f pb/pocketbase pb/pb_data/logs.db*
	@echo "Clean complete!"

clean-deps: ## Remove all dependencies
	@echo "Removing dependencies..."
	@rm -rf sk/node_modules sk/pnpm-lock.yaml
	@cd pb && go clean -modcache
	@echo "Dependencies removed!"

# Utility commands
logs: ## Show PocketBase logs
	@cd pb && tail -f pb_data/logs.db

backup: ## Create a backup of PocketBase data
	@cd pb && cp -r pb_data pb_data_backup_$(shell date +%Y%m%d_%H%M%S)
	@echo "Backup created!"