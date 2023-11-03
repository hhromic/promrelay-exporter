// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package metrics

import (
	"github.com/hhromic/promrelay-exporter/v2/internal/buildinfo"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Namespace is the metrics namespace for the application.
const Namespace = "promrelay"

// BuildInfo is the collector for build information of the application.
//
//nolint:gochecknoglobals
var BuildInfo = promauto.NewGaugeFunc(
	prometheus.GaugeOpts{
		Namespace: Namespace,
		Subsystem: "build",
		Name:      "info",
		Help: "A metric with a constant '1' value labeled by version, goversion, gitcommit, " +
			"gitbranch, builddate from which the application was built.",
		ConstLabels: prometheus.Labels{
			"version":   buildinfo.Version,
			"goversion": buildinfo.GoVersion,
			"gitcommit": buildinfo.GitCommit,
			"gitbranch": buildinfo.GitBranch,
			"builddate": buildinfo.BuildDate,
		},
	},
	func() float64 { return 1 },
)

// RelayInFlightRequests is the collector for the number of relay requests currently being served.
//
//nolint:gochecknoglobals
var RelayInFlightRequests = promauto.NewGauge(
	prometheus.GaugeOpts{
		Namespace:   Namespace,
		Subsystem:   "relay",
		Name:        "in_flight_requests",
		Help:        "Number of relay requests currently being served in the Prometheus relay exporter.",
		ConstLabels: prometheus.Labels{},
	},
)

// RelayRequestsTotal is the collector for the total number of relay requests.
//
//nolint:gochecknoglobals
var RelayRequestsTotal = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Namespace:   Namespace,
		Subsystem:   "relay",
		Name:        "requests_total",
		Help:        "Total number of relay requests in the Prometheus relay exporter.",
		ConstLabels: prometheus.Labels{},
	},
	[]string{"code"},
)

// RelayRequestDuration is the collector for the distribution of relay request durations.
//
//nolint:exhaustruct,gochecknoglobals
var RelayRequestDuration = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Namespace:   Namespace,
		Subsystem:   "relay",
		Name:        "request_duration_seconds",
		Help:        "Distribution of relay request durations in the Prometheus relay exporter.",
		Buckets:     []float64{.1, .2, .4, 1, 3, 8, 20, 60, 120},
		ConstLabels: prometheus.Labels{},
	},
	[]string{"code"},
)

// RelayResponseSize is the collector for the distribution of relay response sizes.
//
//nolint:exhaustruct,gochecknoglobals
var RelayResponseSize = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Namespace:   Namespace,
		Subsystem:   "relay",
		Name:        "response_size_bytes",
		Help:        "Distribution of relay response sizes in the Prometheus relay exporter.",
		Buckets:     prometheus.ExponentialBuckets(100, 10, 8), //nolint:gomnd
		ConstLabels: prometheus.Labels{},
	},
	[]string{"code"},
)
