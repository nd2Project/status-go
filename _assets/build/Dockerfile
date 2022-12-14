# Build status-go in a Go builder container
FROM golang:1.18-alpine as builder

RUN apk add --no-cache make gcc g++ musl-dev linux-headers

ARG build_tags
ARG build_flags

RUN mkdir -p /go/src/github.com/status-im/status-go
WORKDIR /go/src/github.com/status-im/status-go
ADD . .
RUN make statusgo BUILD_TAGS="$build_tags" BUILD_FLAGS="$build_flags"

# Copy the binary to the second image
FROM alpine:latest

LABEL maintainer="support@status.im"
LABEL source="https://github.com/status-im/status-go"
LABEL description="status-go is an underlying part of Status - a browser, messenger, and gateway to a decentralized world."

RUN apk add --no-cache ca-certificates bash libgcc libstdc++
RUN mkdir -p /static/keys

COPY --from=builder /go/src/github.com/status-im/status-go/build/bin/statusd /usr/local/bin/
COPY --from=builder /go/src/github.com/status-im/status-go/static/keys/* /static/keys/

# 30304 is used for Discovery v5
EXPOSE 8080 8545 30303 30303/udp 30304/udp

ENTRYPOINT ["/usr/local/bin/statusd"]
CMD ["--help"]
