variable "TAG" {
  default = "latest"
}

variable "IMAGE_CREATED" {
  default = timestamp()
  type = string
}

variable "_metadata" {
  default = {
    "org.opencontainers.image.authors" = "https://github.com/hhromic"
    "org.opencontainers.image.created" = IMAGE_CREATED
    "org.opencontainers.image.description" = "Simple Prometheus relay exporter written in Go for scraping applications in isolated networks."
    "org.opencontainers.image.documentation" = "https://github.com/hhromic/promrelay-exporter"
    "org.opencontainers.image.licenses" = "Apache-2.0"
    "org.opencontainers.image.source" = "https://github.com/hhromic/promrelay-exporter"
    "org.opencontainers.image.title" = "Prometheus Relay Exporter"
    "org.opencontainers.image.url" = "https://github.com/hhromic/promrelay-exporter"
    "org.opencontainers.image.vendor" = "https://github.com/hhromic"
  }
  type = map(string)
}

target "default" {
  tags = ["ghcr.io/hhromic/promrelay-exporter:${TAG}"]
  annotations = [for k, v in _metadata: "${k}=${v}"]
  labels = _metadata
  attest = [
    {"type" = "provenance", "disabled" = true},
    {"type" = "sbom", "disabled" = true},
  ]
}

target "snapshot" {
  inherits = ["default"]
  args = {
    GORELEASER_EXTRA_ARGS = "--snapshot"
  }
  tags = ["ghcr.io/hhromic/promrelay-exporter:snapshot"]
}
