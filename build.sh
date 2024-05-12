#!/usr/bin/env bash

# 编译类型设置
os="linux darwin windows"
arch="arm64 amd64"
winSuffix=".exe"

# 版本参数配置
program=StreamMusicLyric
version=0.1.1
branch=b
versions=v${version}.${branch}
head=$(git rev-parse --abbrev-ref HEAD)
commit=$(git rev-parse HEAD)
time=$(date -u '+%Y年%m月%d日%H时%M分%S秒')

# 初始化依赖包
go mod tidy
go mod vendor

# 交叉编译 amd64 linux
for GOOS in ${os}; do
  for GOARCH in ${arch}; do
    go env -w GOOS=${GOOS} GOARCH=${GOARCH}
    if [ "${GOOS}" == "windows" ]; then
      filename=${program}-${GOOS}-${GOARCH}-${versions}${winSuffix}
    else
      filename=${program}-${GOOS}-${GOARCH}-${versions}
    fi
    echo "${filename}，开始编译"
    go build -ldflags "\
    -X main.Version=v${version}-${branch} \
    -X main.Branch=${head} \
    -X main.Commit=${commit} \
    -X main.BuildTime=${time} \
    -X main.GOOS=${GOOS} \
    -X main.GOARCH=${GOARCH} \
    " -v -o ${filename} ./
    echo "${filename}，编译结束"
  done
done