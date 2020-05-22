# build用
FROM golang:1.13-alpine as build

WORKDIR /go/app

COPY . .

RUN apk add --no-cache git bash && \
 go build -o go_app

# 本番用
FROM alpine

ENV APP_HOME /app
WORKDIR $APP_HOME

COPY --from=build /go/app/go_app .

RUN addgroup go \
  && adduser -D -G go go \
  && chown -R go:go $APP_HOME

RUN GRPC_HEALTH_PROBE_VERSION=v0.3.1 && \
  wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
  chmod +x /bin/grpc_health_probe

CMD ["./go_app"]
