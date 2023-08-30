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

// RelayRequestDuration is the collector for the distribution of relay request durations.
//
//nolint:exhaustruct,exhaustivestruct,gochecknoglobals
var RelayRequestDuration = promauto.NewHistogram(
	prometheus.HistogramOpts{
		Namespace:   Namespace,
		Subsystem:   "relay",
		Name:        "request_duration_seconds",
		Help:        "Distribution of relay request durations in the Prometheus relay exporter.",
		ConstLabels: prometheus.Labels{},
	},
)

// RelayRequestErrors is the collector for the total number of relay request errors.
//
//nolint:gochecknoglobals
var RelayRequestErrors = promauto.NewCounter(
	prometheus.CounterOpts{
		Namespace:   Namespace,
		Subsystem:   "relay",
		Name:        "request_errors_total",
		Help:        "Total number of relay request errors in the Prometheus relay exporter.",
		ConstLabels: prometheus.Labels{},
	},
)
