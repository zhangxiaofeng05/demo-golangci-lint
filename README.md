## install golangci-lint
https://golangci-lint.run/usage/install/#install-from-source

## usage
### Installation
|Method|Command|
|----|----|
|curl|`sh -c "$(curl -fsSL https://raw.githubusercontent.com/zhangxiaofeng05/demo-golangci-lint/main/install.sh)"`|
|wget|`sh -c "$(wget -O- https://raw.githubusercontent.com/zhangxiaofeng05/demo-golangci-lint/main/install.sh)"`|

### Manual lint
```bash
make lint
```

## QA
Q: no lint specify code  
A: https://golangci-lint.run/usage/false-positives/#nolint-directive

Q: add or remove linters  
A: https://golangci-lint.run/usage/linters/

## reference
1. https://tutorialedge.net/golang/improving-go-workflow-with-git-hooks/
