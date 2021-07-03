# prometheus-relay

A simple [Prometheus](https://prometheus.io/) relay server written in
[Node.js](https://nodejs.org/) for scraping applications in private networks.

The motivating use case for this project can be found [here](use-case.md).

## Building

To build a Docker image for the project:

    docker build -t prometheus-relay .

> **Note:** Ready-to-use images are available in the
> [GitHub Container Registry](https://github.com/users/hhromic/packages/container/package/prometheus-relay).

## Usage

A usage example can be found in the [`example`](example/) directory.

## Code Linting

To lint the codebase of the project using [ESLint](https://eslint.org/) and Docker:

    docker run -t --rm -v $PWD:/data:ro cytopia/eslint .

## License

This project is licensed under the [Apache License Version 2.0](LICENSE).
