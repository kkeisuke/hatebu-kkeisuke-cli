FROM golang:1.13.5-buster

ENV GO111MODULE="on"

WORKDIR /go/src/github.com/kkeisuke/hatebu-kkeisuke-cli

RUN go get \
  golang.org/x/lint/golint \
  # realize の 依存関係の不具合のため、urfave/cli を手動で入れる
  # https://github.com/oxequa/realize/issues/253#issuecomment-532077068
  gopkg.in/urfave/cli.v2@master \
  github.com/oxequa/realize

CMD realize start
