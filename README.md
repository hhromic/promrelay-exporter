# prometheus-relay

A simple [Prometheus](https://prometheus.io/) relay server for scraping
applications behind private networks written in [Node.js](https://nodejs.org/).

## Building

To build a Docker image for the project:

    docker build -t prometheus-relay .

## Usage

To run the Prometheus relay server using Docker:

    docker run --rm -p 8080:8080 prometheus-relay

To run the Prometheus relay server using Docker with a custom port:

    docker run --rm -e PORT=5555 -p 5555:5555 prometheus-relay

## License

This project is licensed under the [Apache License Version 2.0](LICENSE).
