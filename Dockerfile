FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/library/golang:alpine-linuxarm64
RUN apk add --no-cache curl
LABEL authors="JiangWe Leo"

WORKDIR /app
COPY config.yml config.yml
COPY config-log.yml config-log.yml
COPY *linux-arm64* app

RUN chmod +x app
ENTRYPOINT ["./app"]