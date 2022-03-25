# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/ui: run the website ui
.PHONY: run/ui
run/ui:
	@go run ./cmd/ui

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/ui: build the cmd/ui application
.PHONY: build/ui
build/ui:
	@echo 'Building cmd/ui...'
	go build -ldflags='-s' -o=./bin/ui ./cmd/ui

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: tidy dependencies and format
.PHONY: tidy
tidy:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...

## vet: vet all code
.PHONY: vet
vet:
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...

## test: run go tests
.PHONY: test
test:
	@echo 'Running tests...'
	go test -race -vet=off ./...