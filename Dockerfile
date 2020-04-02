# Builder
FROM golang:1.14-alpine as builder

RUN apk update && \
    apk upgrade && \
    apk --update add git gcc

WORKDIR /tmp/src/github.com/cjcjcj/todo

COPY . .

RUN go build -o todo.app main.go

# Distribution
FROM alpine:latest

WORKDIR /app 

COPY --from=builder /tmp/src/github.com/cjcjcj/todo /app

EXPOSE 8080

CMD /app/todo.app
