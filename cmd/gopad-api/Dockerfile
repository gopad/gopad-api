FROM --platform=$BUILDPLATFORM golang:1.25.5-alpine3.21@sha256:b4dbd292a0852331c89dfd64e84d16811f3e3aae4c73c13d026c4d200715aff6 AS builder

RUN apk add --no-cache -U nodejs npm git curl
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

WORKDIR /go/src/gopad
COPY . /go/src/gopad/

RUN --mount=type=cache,target=/go/pkg \
    go mod download -x

ARG TARGETOS
ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg \
    --mount=type=cache,target=/root/.cache/go-build \
    task fe:install fe:generate fe:build && \
    task be:generate build:server GOOS=${TARGETOS} GOARCH=${TARGETARCH}

FROM alpine:3.23@sha256:25109184c71bdad752c8312a8623239686a9a2071e8825f20acb8f2198c3f659

RUN apk add --no-cache ca-certificates mailcap && \
    addgroup -g 1337 gopad && \
    adduser -D -u 1337 -h /var/lib/gopad -G gopad gopad

EXPOSE 8000 8080
VOLUME ["/var/lib/gopad"]
ENTRYPOINT ["/usr/bin/gopad-api"]
CMD ["server"]
HEALTHCHECK CMD ["/usr/bin/gopad-api", "health"]

ENV GOPAD_API_DATABASE_DRIVER="sqlite3"
ENV GOPAD_API_DATABASE_NAME="/var/lib/gopad/database.sqlite3"
ENV GOPAD_API_UPLOAD_DRIVER="file"
ENV GOPAD_API_UPLOAD_PATH="/var/lib/gopad/uploads"

COPY --from=builder /go/src/gopad/bin/gopad-api /usr/bin/gopad-api
COPY --from=builder --chown=1337:1337 /go/src/gopad/pkg/frontend/files /usr/share/gopad
WORKDIR /var/lib/gopad
USER gopad
