## install golangci-lint
https://golangci-lint.run/usage/install/#install-from-source

## usage
1. Add the following content to the makefile
```
git_hook:
	@cp .hooks/* .git/hooks

# golangci-lint
lint:
	@golangci-lint --version
	@golangci-lint run ./...	
```
2. get pre-commit
```shell
wget -P .hooks https://raw.githubusercontent.com/zhangxiaofeng05/demo-golangci-lint/main/.hooks/pre-commit
```
3. get golangci-lint config file
```
wget https://raw.githubusercontent.com/zhangxiaofeng05/demo-golangci-lint/main/.golangci.yml
```
4. Add execution permission to per-commit
```
chmod +x .hooks/pre-commit
```
5. put pre-commit to .git/hooks
```shell
make git_hook
```

ok, let us enjoy it!

## QA
Q: hook was ignored?
```
hint: The '.git/hooks/pre-commit' hook was ignored because it's not set as executable.
hint: You can disable this warning with `git config advice.ignoredHook false`.
```
A: Need to add execute permission
```
chmod +x .hooks/pre-commit
make deps
```
