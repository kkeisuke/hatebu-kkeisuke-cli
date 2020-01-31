FROM golang:1.13.7-buster

WORKDIR /go/src/github.com/kkeisuke/hatebu-kkeisuke-cli

RUN go get \
  golang.org/x/lint/golint

RUN curl -fLo /go/bin/air https://git.io/linux_air \
  && chmod +x /go/bin/air

CMD air
