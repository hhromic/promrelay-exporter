# prometheus-relay-exporter

A simple [Prometheus](https://prometheus.io/) relay exporter written in
[Node.js](https://nodejs.org/) for scraping applications in isolated networks.

This exporter uses the [multi-target exporter pattern](https://prometheus.io/docs/guides/multi-target-exporter/)
described in the Prometheus documentation.

The motivating use case for this project can be found [here](use-case.md).

## Building

To build a Docker image for the project:

    docker build -t prometheus-relay-exporter .

> **Note:** Ready-to-use images are available in the
> [GitHub Container Registry](https://github.com/users/hhromic/packages/container/package/prometheus-relay-exporter).

## Usage

A full example can be found in the [`example`](example/) directory.

## Code Standard

The codebase of this project adheres to the [JavaScript Standard Style](https://standardjs.com/).

To check for compliance (requires a local installation of [Node.js](https://nodejs.org/)):

    npx standard

## License

This project is licensed under the [Apache License Version 2.0](LICENSE).
