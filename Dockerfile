FROM golang:1.19-alpine AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/go-sample-app

COPY . .

RUN go get && go build

FROM alpine:3

ARG BUILD_DATE
ARG BUILD_VCS_REF

ENV TCP_ENABLED=true
ENV UDP_ENABLED=true
ENV LOG_LEVEL=INFO
ENV LOG_FORMAT=plain

ADD docker /app

COPY --from=build_base /tmp/go-sample-app/tod_server /app/tod_server

WORKDIR /app

ENTRYPOINT ["/app/init.sh"]

LABEL org.opencontainers.image.vendor="appleJuiceNETZ" \
      org.opencontainers.image.url="https://applejuicenet.cc" \
      org.opencontainers.image.created=${BUILD_DATE} \
      org.opencontainers.image.revision=${BUILD_VCS_REF} \
      org.opencontainers.image.source="https://github.com/applejuicenetz/tod_server"
