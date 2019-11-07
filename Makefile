test:
	@go test -bench=. -race -coverpkg=. -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out
	@rm coverage.out