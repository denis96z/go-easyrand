.PHONY: fmt
fmt:
	go fmt ./...
	gofumpt -l -w .

.PHONY: test
test:
	go test ./...

.PHONY: deps
deps:
	go mod tidy && go mod verify

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: bootstrap
bootstrap:
	go get mvdan.cc/gofumpt@v0.1.1
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.38.0
