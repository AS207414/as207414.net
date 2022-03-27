# ==================================================================================== #
# VERSIONING VARIABLES
# ==================================================================================== #

git_description = $(shell git describe --always --dirty --tags)
linker_flags = '-s -w -X main.version=${git_description}'

# ==================================================================================== #
# BUILD VARIABLES
# ==================================================================================== #

GOOS = 'linux'
GOARCH = 'amd64'
build_directory = 'bin'
build_outfile = '${build_directory}/as207414_${GOOS}_${GOARCH}'
docker_image_name = "as207414_ui"

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
	@docker run --rm -p 4000:4000 -it ${docker_image_name}

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/ui: build the cmd/ui application
.PHONY: build/ui
build/ui:
	@echo 'Building for ${GOOS}_${GOARCH} to ${build_outfile}'
	@CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -a -installsuffix cgo -ldflags=${linker_flags} -o=${build_outfile} ./cmd/ui

## build/ui/docker: build the cmd/ui docker image
.PHONY: build/ui/docker
build/ui/docker:
	@echo 'Building docker image for ${GOOS}_${GOARCH} to ${docker_image_name}'
	@scripts/build.sh -f ${docker_image_name}

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