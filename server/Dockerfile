# 開発環境
FROM golang:1.11.5

WORKDIR /go/src/api
COPY . .
ENV GO111MODULE=on

RUN go get github.com/pilu/fresh
CMD ["fresh"]