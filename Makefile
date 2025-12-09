DATABASE_URL=postgres://myuser:mypassword@localhost:5432/mesa-mestre?sslmode=disable

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
lint:
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

.PHONY: sqlc-gen
sqlc-gen:
	@sqlc generate -f ./gateway/postgres/sqlc.yaml

.PHONY: migrate-create
migrate-create:
	@read -p "Nome da migração: " name; \
	migrate create -ext sql -dir extension/database/priv/migrations -seq $${name}

.PHONY: migrate-up
migrate-up:
	@migrate -path extension/database/priv/migrations -database "$(DATABASE_URL)" up

.PHONY: migrate-down
migrate-down:
	@migrate -path extension/database/priv/migrations -database "$(DATABASE_URL)" down	

.PHONY: migrate-drop
migrate-drop:
	@migrate -path extension/database/priv/migrations -database "$(DATABASE_URL)" drop
