FROM webhippie/alpine:latest

LABEL maintainer="Thomas Boerger <thomas@webhippie.de>" \
  org.label-schema.name="Gopad API" \
  org.label-schema.vendor="Thomas Boerger" \
  org.label-schema.schema-version="1.0"

EXPOSE 8080 8090
VOLUME ["/var/lib/gopad"]

ENTRYPOINT ["/usr/bin/gopad-api"]
CMD ["server"]

ENV GOPAD_API_UPLOAD_DSN file://var/lib/gopad/

RUN apk add --no-cache ca-certificates mailcap bash

COPY dist/binaries/gopad-api-*-linux-amd64 /usr/bin/
