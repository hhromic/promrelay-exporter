variable "TAG" {
  default = "latest"
}

target "default" {
  tags = ["ghcr.io/hhromic/promrelay-exporter:${TAG}"]
}

target "snapshot" {
  inherits = ["default"]
  args = {
    GORELEASER_EXTRA_ARGS = "--snapshot"
  }
  tags = ["ghcr.io/hhromic/promrelay-exporter:snapshot"]
}
