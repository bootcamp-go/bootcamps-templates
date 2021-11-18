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
	@echo "MysqlRoot Passowrd (if don't have ignore): "; \
    read PASS; \
    curl -s https://raw.githubusercontent.com/bootcamp-go/bootcamps-scripts/main/meli_database.sh | bash  -s create $$PASS

.PHONY: rebuild-database
rebuild-database:
	@echo "MysqlRoot Passowrd (if don't have ignore): "; \
    read PASS; \
    curl -s https://raw.githubusercontent.com/bootcamp-go/bootcamps-scripts/main/meli_database.sh | bash  -s rebuild $$

.PHONY: rebuild-database-with-password	
rebuild-database-with-password:
	@echo "MysqlRoot Passowrd (if don't have ignore): "; \
    read PASS; \
    curl -s https://raw.githubusercontent.com/bootcamp-go/bootcamps-scripts/main/meli_database.sh | bash  -s rebuild ${p}