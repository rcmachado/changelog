FROM golang:alpine AS build

WORKDIR /app

COPY . /app
RUN go build -o changelog .

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/changelog /usr/local/bin/changelog

ENTRYPOINT ["/usr/local/bin/changelog"]
