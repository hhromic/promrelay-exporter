// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"fmt"
	"net/http"
	"net/http/httputil"

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
		ctx := r.Context()

		target, ok := TargetFromContext(ctx)
		if !ok || target == nil {
			err := fmt.Errorf("no target in request context")
			handleErr(w, err, http.StatusInternalServerError)
			return
		}

		rp := &httputil.ReverseProxy{
			ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
				err = fmt.Errorf("target relay error: %w", err)
				handleErr(w, err, http.StatusBadGateway)
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
