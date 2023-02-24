git_hook:
	@cp .hooks/* .git/hooks

# golangci-lint
lint:
	@golangci-lint --version
	@golangci-lint run ./...
