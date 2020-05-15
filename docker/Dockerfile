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

CMD ["./go_app"]
