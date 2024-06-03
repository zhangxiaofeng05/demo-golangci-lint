#!/bin/sh

# 取消 git commit 之前的检查
#if [ ! -d .git ]; then
#	echo "The current directory must be a root Git repository."
#	exit 1
#fi

#if [ -e ".hooks/pre-commit" ]; then
#	echo ".hooks/pre-commit exist!!!"
#	make git_hook
#	echo "finish git_hook, you can enjoy it!"
#	exit 1
#fi

#echo "** start init"
# get pre-commit
#echo "start get pre-commit"
#wget -P .hooks https://raw.githubusercontent.com/zhangxiaofeng05/demo-golangci-lint/main/.hooks/pre-commit
# Add the following content to the makefile
#echo "start appent lint to Makefile"
#curl https://raw.githubusercontent.com/zhangxiaofeng05/demo-golangci-lint/main/Makefile >> Makefile
# get golangci-lint config file
#echo "start get golangci-lint config file"
#wget https://raw.githubusercontent.com/zhangxiaofeng05/demo-golangci-lint/main/.golangci.yml
# Add execution permission to per-commit
#echo "start add execution permission to per-commit"
#chmod +x .hooks/pre-commit
# put pre-commit to .git/hooks
#echo "start put pre-commit to .git/hooks"
#make git_hook

#echo "** ok, let us enjoy it!"

ACTION_INSTALL="install"
ACTION_UPDATE="update"

start() {
  local action=$1
  case "$action" in
  "$ACTION_INSTALL")
    run_install
    ;;
  "$ACTION_UPDATE")
    run_update
    ;;
  *)
    echo "missing action, default install"
    run_install
    ;;
  esac
}

run_install() {
  echo "start install"
  # 检查 golangci-lint 是否存在
  if
    ! command -v golangci-lint &
    >/dev/null
  then
    # echo "golangci-lint 未安装或未添加到系统 PATH 中"
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
  fi
  golangci-lint --version
  # 检查 .golangci.yml
  if [ -e ".golangci.yml" ]; then
    echo ".golangci.yml 文件存在，需要更新可执行更新脚本,忽略下载"
  else
    wget https://raw.githubusercontent.com/zhangxiaofeng05/demo-golangci-lint/main/.golangci.yml
  fi
  echo "手动 lint,不再使用git hooks, 将 golangci-lint run ./... 写入到 Makefile中"
  echo "end install"
}

run_update() {
  echo "start update"
  # echo "golangci-lint 未安装或未添加到系统 PATH 中"
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
  golangci-lint --version
  rm -rf .golangci.yml
  wget https://raw.githubusercontent.com/zhangxiaofeng05/demo-golangci-lint/main/.golangci.yml
  echo "end update"
}

# 检查是否提供了参数
# if [ $# -ne 1 ]; then
# echo "Usage: $0 {install|update}"
# exit 1
# fi

# 将命令行参数传递给函数
start "$1"
