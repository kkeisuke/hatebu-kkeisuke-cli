FROM golang:1.13.5-buster

ENV GO111MODULE="on"

WORKDIR /go/src/github.com/kkeisuke/hatebu-kkeisuke-cli

RUN go get \
  golang.org/x/lint/golint

RUN curl -fLo /go/bin/air https://raw.githubusercontent.com/cosmtrek/air/master/bin/linux/air
RUN chmod +x /go/bin/air

CMD air
