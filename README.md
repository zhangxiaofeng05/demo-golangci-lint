## 使用
### curl 或者 wget 二选一
|Method|Command|
|----|----|
|curl|`sh -c "$(curl -fsSL https://raw.githubusercontent.com/zhangxiaofeng05/demo-golangci-lint/main/install.sh)"`|
|wget|`sh -c "$(wget -O- https://raw.githubusercontent.com/zhangxiaofeng05/demo-golangci-lint/main/install.sh)"`|

shell做的事：安装`golangci-lint`，下载`.golangci.yml`。

手动执行lint的命令
```bash
golangci-lint run ./...
```
如果使用Makefile，将下边的命令添加到Makefile中
```makefile
# lint
lint:
	@golangci-lint --version
	@golangci-lint run ./...
```

## 依赖 golangci-lint
https://golangci-lint.run/usage/install/#install-from-source

## QA
Q: no lint specify code  
A: https://golangci-lint.run/usage/false-positives/#nolint-directive

Q: add or remove linters  
A: https://golangci-lint.run/usage/linters/

## reference
1. https://tutorialedge.net/golang/improving-go-workflow-with-git-hooks/
