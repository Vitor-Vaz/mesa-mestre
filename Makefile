run:
	go run ./api/main.go

.PHONY: run

.PHONY: install-linters
install-linters:
	@echo "📦 Instalando todas as dependências..."
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install golang.org/x/lint/golint@latest
	@go install github.com/kisielk/errcheck@latest
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@go install github.com/securego/gosec/v2/cmd/gosec@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "✅ Todas as dependências instaladas!"

.PHONY: lint
lint: ## Executa todas as verificações de lint
	@echo "🎨 Formatando código..."
	@goimports -w .
	@gofmt -w -s .
	@echo "⚠️  Verificando erros não tratados..."
	@errcheck ./...
	@echo "🔒 Executando gosec..."
	@gosec ./...
	@echo "🔍 Executando golangci-lint..."
	@golangci-lint run ./...
	@echo "🔧 Corrigindo problemas de lint..."
	@golangci-lint run --fix ./...
	@echo "🔍 Executando go vet..."
	@go vet ./...
	@echo "🔍 Executando staticcheck..."
	@staticcheck ./...



