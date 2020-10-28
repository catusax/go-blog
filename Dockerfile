FROM golang:alpine as builder

RUN apk --no-cache add git \
        && cd / \
        && git clone https://github.com/coolrc136/go-blog.git \
        && cd go-blog \
        && CGO_ENABLED=0 GOOS=linux go build

FROM alpine

COPY --from=0 /go-blog/blog /blog/
COPY --from=0 /go-blog/static /blog/static
COPY --from=0 /go-blog/config /blog/config
CMD cd /blog && ./blog
