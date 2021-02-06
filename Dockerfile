FROM golang:1.13-alpine 

ENV GO111MODULE=on

RUN apk add --no-cache git libgit2-dev alpine-sdk

WORKDIR /go/src/github.com/jekabolt/quotes

# https://divan.github.io/posts/go_get_private/
# COPY .gitconfig /root/.gitconfig
COPY go.mod .
COPY go.sum .
# install dependencies
RUN go mod download

COPY ./cmd/ ./cmd/
COPY ./signature/ ./signature/
COPY ./server/ ./server/

RUN go build -o ./bin/quotes-server ./cmd/

FROM alpine:latest

WORKDIR /go/src/github.com/jekabolt/quotes
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
# RUN apk add --no-cache git libgit2-dev alpine-sdk
RUN apk --no-cache add curl

COPY --from=0 /go/src/github.com/jekabolt/quotes .

CMD ["/go/src/github.com/jekabolt/quotes/bin/quotes-server"]
