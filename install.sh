#!/bin/sh

# Add the following content to the makefile
echo "start appent lint to Makefile"
curl https://raw.githubusercontent.com/zhangxiaofeng05/demo-golangci-lint/main/Makefile >> Makefile
# get pre-commit
echo "start get pre-commit"
wget -P .hooks https://raw.githubusercontent.com/zhangxiaofeng05/demo-golangci-lint/main/.hooks/pre-commit
# get golangci-lint config file
echo "start get golangci-lint config file"
wget https://raw.githubusercontent.com/zhangxiaofeng05/demo-golangci-lint/main/.golangci.yml
# Add execution permission to per-commit
echo "start add execution permission to per-commit"
chmod +x .hooks/pre-commit
# put pre-commit to .git/hooks
echo "start put pre-commit to .git/hooks"
make git_hook

echo "ok, let us enjoy it!"