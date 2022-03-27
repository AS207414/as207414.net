# ==================================================================================== #
# BUILD VARIABLES
# ==================================================================================== #

GOOS = 'linux'
GOARCH = 'amd64'
DOCKER_IMAGE_NAME = "as207414_ui"

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

## run/ui/docker: run the website ui from docker image
.PHONY: run/ui/docker
run/ui/docker:
	@docker run --rm -p 4000:4000 -it ${DOCKER_IMAGE_NAME}

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/ui: build the cmd/ui application
.PHONY: build/ui
build/ui:
	@echo 'Building for ${GOOS}_${GOARCH}'
	@CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} scripts/build-go.sh -s ui

## build/ui/docker: build the cmd/ui docker image
.PHONY: build/ui/docker
build/ui/docker:
	@echo 'Building docker image for ${GOOS}_${GOARCH} to ${DOCKER_IMAGE_NAME}'
	@scripts/build-docker.sh -f ${DOCKER_IMAGE_NAME}

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