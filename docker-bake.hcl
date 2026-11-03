variable "TAG" {
  default = "latest"
}

target "default" {
  tags = ["ghcr.io/hhromic/promrelay-exporter:${TAG}"]
  annotations = [
    "org.opencontainers.image.source=https://github.com/hhromic/promrelay-exporter",
    "org.opencontainers.image.description=Simple Prometheus relay exporter written in Go for scraping applications in isolated networks.",
    "org.opencontainers.image.licenses=Apache-2.0",
  ]
}

target "snapshot" {
  inherits = ["default"]
  args = {
    GORELEASER_EXTRA_ARGS = "--snapshot"
  }
  tags = ["ghcr.io/hhromic/promrelay-exporter:snapshot"]
}
