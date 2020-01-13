FROM golang:alpine

WORKDIR /app

COPY . /app
RUN go build -o changelog .

FROM alpine:latest

WORKDIR /app
COPY --from=0 /app/changelog /usr/local/bin/changelog

CMD ["/usr/local/bin/changelog"]
