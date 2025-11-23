# Static Example

> [!NOTE]
> The example in this directory assumes that you have Docker running
> in [swarm mode](https://docs.docker.com/engine/swarm/).

To run the example, first deploy the centralized Prometheus stack:
```
docker config create metrics-prometheus-yaml configs/prometheus.yaml
docker stack deploy -c stacks/metrics.yaml metrics
```

The above will initialize a Prometheus instance running on port `9090`.

In addition, this stack also creates a `metrics_default` overlay network.

> [!NOTE]
> This example Prometheus instance uses static targets for simplicity.
> However, [automatic discovery](../autodiscovery/) of containers can be also implemented.

At this point, you can browse the running [Prometheus UI](http://localhost:9090).

If you navigate to the [targets status page](http://localhost:9090/targets), you
should be able to find a `node` job with a non-operational endpoint (DOWN).

Now, to deploy the example application stack:
```
docker stack deploy -c stacks/app.yaml app
```

This stack will deploy two containers:

* A Prometheus Node Exporter container with some metrics to scrape (`node`).
* A Prometheus Relay Exporter container (`promrelay`).

You can now verify that Prometheus is able to scrape the metrics exposed by the
Node Exporter container, without having direct access to it.

Finally, to remove everything done in this example:
```
docker stack rm app
docker stack rm metrics
docker config rm metrics-prometheus-yaml
docker volume rm metrics_data
```
