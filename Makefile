.PHONY: run
run:
	@go run cmd/alpha.go

tests:
	@go test -cover -count 1 -parallel 8 ./...