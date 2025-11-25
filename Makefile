run:
	go run ./app/main.go

.PHONY: run

.PHONY: install-linters
install-linters:
	@echo "📦 Installing all dependencies..."
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install golang.org/x/lint/golint@latest
	@go install github.com/kisielk/errcheck@latest
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@go install github.com/securego/gosec/v2/cmd/gosec@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "✅ All dependencies installed!"

.PHONY: lint
lint: ## Execute linters and formatters
	@echo "🎨 Formatting code..."
	@goimports -w .
	@gofmt -w -s .
	@echo "⚠️  Checking for unchecked errors..."
	@errcheck ./...
	@echo "🔒 Running gosec..."
	@gosec ./...
	@echo "🔍 Running golangci-lint..."
	@golangci-lint run ./...
	@echo "🔧 Fixing lint issues..."
	@golangci-lint run --fix ./...
	@echo "🔍 Running go vet..."
	@go vet ./...
	@echo "🔍 Running staticcheck..."
	@staticcheck ./...



