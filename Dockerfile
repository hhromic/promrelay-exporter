# Start a new stage for building the application
FROM golang:1.23.5 AS builder

# Download and install GoReleaser
RUN TGZ_URL=https://github.com/goreleaser/goreleaser/releases/download/v2.6.1/goreleaser_Linux_x86_64.tar.gz \
    && curl --proto '=https' --tlsv1.2 -L "${TGZ_URL}" | tar zxf - -C /usr/bin goreleaser

# Set a well-known building directory
WORKDIR /build

# Download and verify application dependencies
COPY go.mod go.sum ./
RUN go mod download \
    && go mod verify

# Copy application sources and build the application
COPY . .
ARG GORELEASER_EXTRA_ARGS
RUN GOTOOLCHAIN=local \
    CGO_ENABLED=0 \
    goreleaser build --clean --single-target ${GORELEASER_EXTRA_ARGS}

# Start a new stage for the final application image
FROM cgr.dev/chainguard/static:latest AS final

# Configure image labels
LABEL org.opencontainers.image.source=https://github.com/hhromic/promrelay-exporter \
      org.opencontainers.image.description="Simple Prometheus relay exporter written in Go for scraping applications in isolated networks." \
      org.opencontainers.image.licenses=Apache-2.0

# Configure default entrypoint and exposed port of the application
ENTRYPOINT ["/promrelay-exporter"]
EXPOSE 9878

# Copy application binary
COPY --from=builder /build/dist/promrelay-exporter /promrelay-exporter
