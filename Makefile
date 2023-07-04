.PHONY: build
build:
	@go build -o alpha cmd/alpha.go

.PHONY: run
run:
	@go run cmd/alpha.go

.PHONY: tests
tests:
	@go test -cover -count 1 -parallel 8 ./...