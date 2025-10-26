.PHONY: dev setup clean

setup:
	@cd server && go mod tidy
	@cd client && npm install
	@cd client && npm install --save-dev concurrently

dev:
	@cd client && npx concurrently \
		--names "backend,frontend" \
		--prefix-colors "blue,green" \
		"cd ../server && go run cmd/main.go" \
		"cd . && npm run dev"

clean:
	@cd server && go clean
	@cd client && rm -rf node_modules dist