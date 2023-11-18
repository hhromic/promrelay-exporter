// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"net/http"

	"github.com/hhromic/promrelay-exporter/v2/internal/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	// PatternMetricsHandler is the path pattern to use for the metrics handler.
	PatternMetricsHandler = "/metrics"
	// PatternRelayHandler is the path pattern to use for the relay handler.
	PatternRelayHandler = "/relay"
)

// NewServeMux creates a top-level request multiplexer for the application.
func NewServeMux() *http.ServeMux {
	rhandler := promhttp.InstrumentHandlerInFlight(
		metrics.RelayInFlightRequests,
		promhttp.InstrumentHandlerDuration(
			metrics.RelayRequestDuration,
			promhttp.InstrumentHandlerCounter(
				metrics.RelayRequestsTotal,
				promhttp.InstrumentHandlerResponseSize(
					metrics.RelayResponseSize,
					RelayHandler(),
				),
			),
		),
	)

	m := http.NewServeMux()
	m.Handle(PatternMetricsHandler, promhttp.Handler())
	m.Handle(PatternRelayHandler, rhandler)

	return m
}
