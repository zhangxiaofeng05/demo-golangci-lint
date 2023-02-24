## install golangci-lint
https://golangci-lint.run/usage/install/#install-from-source

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
