FROM caddy:2-builder AS builder

RUN go env -w GOPROXY=https://goproxy.io,direct \
    && xcaddy build \
    --with github.com/caddy-dns/cloudflare

FROM caddy:2-alpine

COPY --from=builder /usr/bin/caddy /usr/bin/caddy
