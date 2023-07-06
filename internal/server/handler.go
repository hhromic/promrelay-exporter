// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/hhromic/promrelay-exporter/internal/metrics"
	"golang.org/x/exp/slog"
)

// RelayHandler is an http.Handler for target relay requests.
func RelayHandler() http.Handler {
	handleErr := func(w http.ResponseWriter, err error, status int) {
		http.Error(w, err.Error(), status)
		slog.Error("relay handler error", "err", err, "status", status)
		metrics.RelayRequestErrors.Add(1)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var target *url.URL

		s := time.Now()
		defer func() {
			d := time.Since(s)
			slog.Debug("relay request completed", "target", target, "duration", d)
			metrics.RelayRequestDuration.Observe(d.Seconds())
		}()

		query := r.URL.Query()

		if len(query["target"]) != 1 || query["target"][0] == "" {
			handleErr(w,
				fmt.Errorf("'target' parameter is missing or is specified multiple times"),
				http.StatusBadRequest,
			)
			return
		}

		target, err := url.ParseRequestURI(query["target"][0])
		if err != nil {
			handleErr(w,
				fmt.Errorf("'target' parameter is not a valid URL: %w", err),
				http.StatusBadRequest,
			)
			return
		}

		rp := &httputil.ReverseProxy{
			ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
				handleErr(w,
					fmt.Errorf("target relay error: %w", err),
					http.StatusBadGateway,
				)
			},
			Rewrite: func(r *httputil.ProxyRequest) {
				r.SetXForwarded()
				r.Out.URL = target
				r.Out.Host = target.Host
			},
		}

		rp.ServeHTTP(w, r)
	})
}
