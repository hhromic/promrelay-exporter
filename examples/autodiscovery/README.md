# Automatic Discovery Example

> **Note:** The example in this directory assumes that you have Docker running
> in [swarm mode](https://docs.docker.com/engine/swarm/).

To run the example, first deploy the centralised Prometheus stack:
```
docker config create metrics-prometheus-yaml configs/prometheus.yaml
docker stack deploy -c stacks/metrics.yaml metrics
```

The above will initialise a Prometheus instance running on port `9090`.

In addition, this stack also creates a `metrics_default` overlay network.

At this point, you can browse the running [Prometheus UI](http://localhost:9090).

Now, to deploy the example application stack:
```
docker stack deploy -c stacks/app.yaml app
```

This stack will deploy four containers:

* Three Prometheus Node Exporter containers with some metrics to scrape (`myapp`).
  * Note: Metrics scraping is configured using deploy labels in the application service(s).
* A Prometheus Relay Exporter container (`promrelay`).

After deploying, and after a short time, the configured Docker Swarm autodiscovery in Prometheus
will find the deployed service containers and autoconfigure metrics scraping using the labels.

> **Note:** The Docker Swarm autodiscovery defaults to refreshing data every `60s`.

You can now verify that Prometheus is able to scrape the metrics exposed by the
Node Exporter containers, without having direct access to it.

Finally, to remove everything done in this example:
```
docker stack rm app
docker stack rm metrics
docker config rm metrics-prometheus-yaml
```

## Configuration Labels

In this example, the usage of a relay exporter is optional and configurable.

To configure the discovered scrape targets, the following labels are available:

* `metrics.enable`: Set to `true` or `false` to enable/disable metrics collection.
* `metrics.path`: (Optional) Set a custom metrics path for target scraping.
* `metrics.port`: (Optional) set a custom port for target scraping.
* `metrics.scheme`: (Optional) Set the protocol scheme for target scraping.
* `metrics.relay.host`: (Optional) Set the host of the relay exporter to use.
* `metrics.relay.port`: (Optional) Set the port of the relay exporter to use.
* `metrics.relay.scheme`: (Optional) Set the protocol scheme for the relay exporter.

> **Note:** The relay exporter is only used when both `metrics.relay.host` and
> `metrics.relay.port` labels are specified.

These labels must be attached at the service level, i.e. under the `deploy` key
in the YAML stack configuration.
