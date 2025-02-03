# Prometheus Relay Exporter

Simple [Prometheus](https://prometheus.io/) relay exporter written in [Go](https://go.dev/) for
scraping applications in isolated networks.

This exporter uses the [multi-target exporter pattern](https://prometheus.io/docs/guides/multi-target-exporter/)
described in the Prometheus documentation.

The motivating use case for this project can be found [here](use-case.md).

## Usage

Usage examples can be found in the [`examples/`](examples/) directory.

## Building

To build a release Docker image, use [Docker Build Bake](https://docs.docker.com/build/bake/):
```
git checkout vX.Y.Z
TAG=vX.Y.Z docker buildx bake
```

> [!NOTE]
> Ready-to-use images are available in the
> [GitHub Container Registry](https://github.com/users/hhromic/packages/container/package/promrelay-exporter).

To build a snapshot locally Using [GoReleaser](https://goreleaser.com/):
```
goreleaser build --clean --single-target --snapshot
```

## License

This project is licensed under the [Apache License Version 2.0](LICENSE).
