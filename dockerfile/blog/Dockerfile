FROM golang:alpine as builder

WORKDIR /go-blog/
COPY . .
RUN CGO_ENABLED=0  GOPROXY=https://goproxy.io,direct GOOS=linux go build

FROM alpine

WORKDIR /blog/
COPY --from=0 /go-blog/blog .
COPY static ./static
COPY config ./config
ENTRYPOINT /blog/blog
