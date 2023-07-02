# Prometheus Relay Exporter

A simple [Prometheus](https://prometheus.io/) relay exporter written in [Go](https://go.dev/).

This exporter uses the [multi-target exporter pattern](https://prometheus.io/docs/guides/multi-target-exporter/)
described in the Prometheus documentation.

The motivating use case for this project can be found [here](use-case.md).

## Usage

Usage examples can be found in the [`examples`](examples/) directory.

## Building

To build a Docker image for the project:
```
docker buildx build -t promrelay-exporter .
```

> **Note:** Ready-to-use images are available in the
> [GitHub Container Registry](https://github.com/users/hhromic/packages/container/package/promrelay-exporter).

## License

This project is licensed under the [Apache License Version 2.0](LICENSE).
