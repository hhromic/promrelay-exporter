# Start a new stage for building the application
FROM golang:1.25.6 AS builder

# Download and install GoReleaser
ADD https://github.com/goreleaser/goreleaser/releases/download/v2.13.3/goreleaser_Linux_x86_64.tar.gz goreleaser.tar.gz
RUN tar zxf goreleaser.tar.gz --one-top-level=/usr/bin/ goreleaser

# Set a well-known building directory
WORKDIR /build

# Download and verify application dependencies
COPY go.mod go.sum ./
RUN go mod download \
    && go mod verify

# Copy application sources and build the application
COPY . .
ARG GORELEASER_EXTRA_ARGS
RUN CGO_ENABLED=0 \
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
