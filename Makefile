# Image URL to use all building/pushing image targets
SHELL := /bin/bash

APP_NAME ?= awi-infra-guard
BIN_NAME ?= ${APP_NAME}

OS ?= $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH ?= $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))

IMG ?= {CONTAINER_IMAGE}

.PHONY: all
all: generate mocks

.PHONY: generate
generate:
	cd grpc && rm -rf go/infrapb/* js/* ts/*
	cd grpc && protoc --proto_path=proto \
	 --go_out=go/infrapb --go_opt=paths=source_relative \
	 --go-grpc_out=go/infrapb --go-grpc_opt=paths=source_relative \
	 --js_out=import_style=commonjs:js --grpc-web_out=import_style=commonjs,mode=grpcwebtext:js \
	 --grpc-web_out=import_style=typescript,mode=grpcwebtext:ts \
	 proto/*.proto

.PHONY: tools
	go install github.com/vektra/mockery/v2@latest

.PHONY: mocks
mocks:
	rm -rf mocks/*
	mockery --with-expecter --all

.PHONY: run
run: ## Run infra-sdk grpc server
	go run main.go

# Build the docker image
.PHONY: docker-build
docker-build: test
	docker build . -t ${IMG}

# Push the docker image
.PHONY: docker-push
docker-push:
	docker push ${IMG}

.PHONY: build
build: ## Build awi-infra-guard grpc server 
	CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} go build -ldflags="-w -s" -o bin/${BIN_NAME}

.PHONY: build_csp_connector
build_csp_connector:
	CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} go build -ldflags="-w -s" -o bin/cspConnector cmd/cspConnector.go

