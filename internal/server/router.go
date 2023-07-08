// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	// MetricsHandlerPattern is the path pattern to use for the metrics handler.
	MetricsHandlerPattern = "/metrics"
	// RelayHandlerPattern is the path pattern to use for the relay handler.
	RelayHandlerPattern = "/relay"
)

// NewRouter creates a top-level http.Handler router for the application.
func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Mount(MetricsHandlerPattern, promhttp.Handler())
	r.Mount(RelayHandlerPattern, RelayHandler())

	return r
}
