# hook
git_hook:
	@cp .hooks/* .git/hooks

# lint
lint:
	@golangci-lint --version
	@golangci-lint run ./...
