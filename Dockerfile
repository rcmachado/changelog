FROM golang:alpine AS build

WORKDIR /app

COPY . /app
RUN go build -o changelog .

FROM alpine:latest

RUN apk add --no-cache git openssh

WORKDIR /app
COPY --from=build /app/changelog /usr/local/bin/changelog

ENTRYPOINT ["/usr/local/bin/changelog"]
