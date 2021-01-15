.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	go test ./...

.PHONY: deps
deps:
	go mod tidy && go mod vendor && go mod verify
