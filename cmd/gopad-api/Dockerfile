FROM --platform=$BUILDPLATFORM debian:13-slim@sha256:d7e12182ce18b85b93007c1dedf31f2d29e01ccf3182cc4017c709b6259bc132 AS builder
WORKDIR /src

RUN apt-get update && \
    apt-get -y --no-install-recommends install ca-certificates curl git build-essential && \
    rm -rf /var/lib/apt/lists/*

SHELL ["/bin/bash", "-o", "pipefail", "-c"]
ENV PATH="/root/.local/bin:/mise/shims:$PATH" \
    MISE_DATA_DIR="/mise" \
    MISE_CONFIG_DIR="/mise" \
    MISE_CACHE_DIR="/mise/cache" \
    MISE_INSTALL_PATH="/usr/local/bin/mise"

RUN curl --proto '=https' --proto-redir '=https' \
    --fail --show-error --silent --location https://mise.run | sh

COPY mise.toml ./
RUN --mount=type=cache,target=/mise/cache \
   mise trust && mise install

COPY . /src/
RUN --mount=type=cache,target=/root/go/pkg \
    go mod download -x

ARG TARGETOS
ARG TARGETARCH

RUN --mount=type=cache,target=/root/go/pkg \
    --mount=type=cache,target=/root/.cache/go-build \
    task be:generate build:server GOOS=${TARGETOS} GOARCH=${TARGETARCH}

FROM alpine:3.24@sha256:28bd5fe8b56d1bd048e5babf5b10710ebe0bae67db86916198a6eec434943f8b

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

COPY --from=builder /src/bin/gopad-api /usr/bin/gopad-api
COPY --from=builder --chown=1337:1337 /src/pkg/frontend/files /usr/share/gopad
WORKDIR /var/lib/gopad
USER gopad
