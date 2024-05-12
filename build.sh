#!/usr/bin/env bash

# 版本参数配置
prog=StreamMusicLyric
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
GOOS=linux
GOARCH=amd64
go env -w GOOS=${GOOS} GOARCH=${GOARCH}
go build -ldflags "\
-X main.Version=v${version}-${branch} \
-X main.Branch=${head} \
-X main.Commit=${commit} \
-X main.BuildTime=${time} \
-X main.GOOS=${GOOS} \
-X main.GOARCH=${GOARCH} \
" -v -o ${prog}-${GOOS}-${GOARCH}-${versions} ./

# 交叉编译 amd64 linux
GOOS=linux
GOARCH=arm64
go env -w GOOS=${GOOS} GOARCH=${GOARCH}
go build -ldflags "\
-X main.Version=v${version}-${branch} \
-X main.Branch=${head} \
-X main.Commit=${commit} \
-X main.BuildTime=${time} \
-X main.GOOS=${GOOS} \
-X main.GOARCH=${GOARCH} \
" -v -o ${prog}-${GOOS}-${GOARCH}-${versions} ./

# 交叉编译 arm64 darwin
GOOS=darwin
GOARCH=arm64
go env -w GOOS=${GOOS} GOARCH=${GOARCH}
go build -ldflags "\
-X main.Version=v${version}-${branch} \
-X main.Branch=${head} \
-X main.Commit=${commit} \
-X main.BuildTime=${time} \
-X main.GOOS=${GOOS} \
-X main.GOARCH=${GOARCH} \
" -v -o ${prog}-${GOOS}-${GOARCH}-${versions} ./

# 交叉编译 amd64 windows
GOOS=windows
GOARCH=amd64
go env -w GOOS=${GOOS} GOARCH=${GOARCH}
go build -ldflags "\
-X main.Version=v${version}-${branch} \
-X main.Branch=${head} \
-X main.Commit=${commit} \
-X main.BuildTime=${time} \
-X main.GOOS=${GOOS} \
-X main.GOARCH=${GOARCH} \
" -v -o ${prog}-${GOOS}-${GOARCH}-${versions}.exe ./
