// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"net/http"
	"strings"
	"time"

	"github.com/hhromic/promrelay-exporter/internal/metrics"
	"golang.org/x/exp/slog"
)

// RelayRequestTimer is a middleware to measure relay request durations.
func RelayRequestTimer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := time.Now()

		defer func() {
			d := time.Since(s)
			t := strings.Join(r.URL.Query()["target"], " ")
			slog.Debug("relay request completed", "target", t, "duration", d)
			metrics.RelayRequestDuration.Observe(d.Seconds())
		}()

		next.ServeHTTP(w, r)
	})
}
