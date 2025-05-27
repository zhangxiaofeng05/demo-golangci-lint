#!/bin/sh

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
    go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
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
  go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
  golangci-lint --version
  rm -rf .golangci.yml
  wget https://raw.githubusercontent.com/zhangxiaofeng05/demo-golangci-lint/main/.golangci.yml
  echo "end update"
}

# 将命令行参数传递给函数
start "$1"
