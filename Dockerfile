FROM golang:1.17-alpine as server-build

WORKDIR /todo-list
COPY . .
RUN apk upgrade --update && \
    apk --no-cache add git