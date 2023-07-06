ifneq ("$(wildcard .env)","")
  $(info using .env)
  include .env
endif

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

# ==================================================================================== #
# DEV
# ==================================================================================== #

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	@go run ./cmd/api/...

current_time = $(shell date "+%Y-%m-%dT%H:%M:%S%z")
git_description = $(shell git describe --always --dirty --tags --long)
linker_flags = '-s -X main.buildTime=${current_time} -X main.version=${git_description}'

## build/api: build the cmd/api application
.PHONY: build/api
build/api:
	@echo 'Building cmd/api...'
	go build -ldflags=${linker_flags} -o=./bin/api ./cmd/api/main.go


.PHONY: build/consumer
build/consumer:
	go build -ldflags=${linker_flags} -o=./bin/consumer ./cmd/consumer/main.go

PHONY: build/apps
build/apps: build/api build/consumer

# ==================================================================================== #
# DOCKER
# ==================================================================================== #

## docker/setup: setup local environment for development
.PHONY: docker/setup
docker/setup:
	@docker-compose build

## docker/up: start the local stack in background
.PHONY: docker/up
docker/up:
	@docker-compose up -d

## docker/down: stop docker
.PHONY: docker/down
docker/down:
	@docker-compose down

## docker/build/dev/api: build the cmd/api docker image
.PHONY: docker/build/api
docker/build/dev/api:
	@echo 'Building api docker image...'
	docker build -t sword-api -f .setup/build/Dockerfile

## docker/run/api: run the cmd/api docker image THIS COMMAND IS NOT WORKING
.PHONY: docker/run/api
docker/run/api:
	@docker run --rm -p 4600:4600 --network journey-negotiation_default journey-negotiation-bff

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## test: test all code with coverage
.PHONY: test
test:
	go test -race -vet=off -coverpkg ./internal/... -v -coverprofile=cover.out  ./internal/...
	go tool cover -html=cover.out



## audit: tidy dependencies, format and vet all code
.PHONY: format
audit:
	@echo 'Formatting code...'
	go fmt ./...

## tidy: tidy dependencies
.PHONY: tidy
tidy:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify