.PHONY: build
build:
	CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} go build -ldflags="-w -s" -o bin/cspConnector cmd/cspConnector.go
