PACKAGES_PATH = $(shell go list -f '{{ .Dir }}' ./...)

.PHONY: all
all: ensure-deps fmt test

.PHONY: ensure-deps
ensure-deps:
	@echo "=> Syncing dependencies with go mod tidy"
	@go mod tidy

.PHONY: fmt
fmt:
	@echo "=> Executing go fmt"
	@go fmt ./...

.PHONY: test
test:
	@echo "=> Running tests"
	@go test ./... -covermode=atomic -coverpkg=./... -count=1 -race

.PHONY: test-cover
test-cover:
	@echo "=> Running tests and generating report"
	@go test ./... -covermode=atomic -coverprofile=/tmp/coverage.out -coverpkg=./... -count=1
	@go tool cover -html=/tmp/coverage.out

.PHONY: start
start:
	@go run cmd/server/main.go

.PHONY: build-database
build-database:
	@mysql -u $(u) -p$(p) -e "CREATE DATABASE IF NOT EXISTS melisprint"
	@mysql -u $(u) -p$(p) melisprint < db.sql
	@mysql -u $(u) -p$(p) -e "CREATE USER 'meli_sprint_user'@'localhost' IDENTIFIED BY 'Meli_Sprint#123'"
	@mysql -u $(u) -p$(p) -e "GRANT ALL PRIVILEGES ON * . * TO 'meli_sprint_user'@'localhost'"
	@mysql -u $(u) -p$(p) -e "FLUSH PRIVILEGES"