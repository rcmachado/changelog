FROM golang:latest

ENV PROJ_PATH=/go/src/github.com/rcmachado/changelog
RUN go get -v -u github.com/golang/dep/cmd/dep
RUN mkdir -p $PROJ_PATH
WORKDIR $PROJ_PATH

COPY Gopkg.* $PROJ_PATH/
RUN dep ensure -vendor-only

COPY . $PROJ_PATH
RUN go build -o changelog .

FROM alpine:latest

WORKDIR /app
RUN mkdir /lib64 && \
    ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY --from=0 /go/src/github.com/rcmachado/changelog/changelog /usr/local/bin/changelog

CMD ["/usr/local/bin/changelog"]
