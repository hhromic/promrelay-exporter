// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/hhromic/promrelay-exporter/v2/internal/metrics"
	"golang.org/x/exp/slog"
)

const (
	// QueryParamTarget is the request query parameter used for providing the relay target.
	QueryParamTarget = "target"
)

const (
	// ResponseHeaderTimeout is the maximum time to wait for reading an HTTP response header.
	ResponseHeaderTimeout = 60 * time.Second
)

// Errors used by handlers in the server package.
var (
	// ErrQueryParamMissing is returned when a request query parameter is missing.
	ErrQueryParamMissing = errors.New("missing query parameter")
)

// RelayHandler is an http.Handler for target relay requests.
func RelayHandler() http.Handler {
	handleErr := func(w http.ResponseWriter, err error, status int) {
		http.Error(w, err.Error(), status)
		slog.Error("relay handler error", "err", err, "status", status)
		metrics.RelayRequestErrors.Add(1)
	}

	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.ResponseHeaderTimeout = ResponseHeaderTimeout

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var target *url.URL

		s := time.Now()
		defer func() {
			d := time.Since(s)
			slog.Debug("relay request completed", "target", target, "duration", d)
			metrics.RelayRequestDuration.Observe(d.Seconds())
		}()

		q := r.URL.Query()
		tgt := q.Get(QueryParamTarget)
		if tgt == "" {
			handleErr(w,
				fmt.Errorf("%w: %q", ErrQueryParamMissing, QueryParamTarget),
				http.StatusBadRequest,
			)
			return
		}

		target, err := url.ParseRequestURI(tgt)
		if err != nil {
			handleErr(w,
				fmt.Errorf("target is not a valid URL: %w", err),
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
			Transport: transport,
		}

		rp.ServeHTTP(w, r)
	})
}
