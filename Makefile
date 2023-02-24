deps:
	@cp .hooks/* .git/hooks

# golangci-lint
lint:
	# how to install: https://golangci-lint.run/usage/install/#install-from-source
	golangci-lint run ./...
