.PHONY: dev setup clean

setup:
	@cd server && go mod tidy
	@cd client && npm install
	@cd client && npm install --save-dev concurrently

dev:
	@echo "Starting backend server..."
	@cd server && go run cmd/main.go &
	@echo "Waiting for backend to start..."
	@sleep 2
	@echo "Starting frontend server..."
	@cd client && npm run dev

clean:
	@cd server && go clean
	@cd client && rm -rf node_modules dist