// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/hhromic/promrelay-exporter/v2/internal/metrics"
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
	ErrQueryParamMissing = errors.New("query parameter missing")
)

// RelayHandler is an [http.Handler] for target relay requests.
func RelayHandler() http.Handler {
	transport := http.DefaultTransport.(*http.Transport).Clone() //nolint:forcetypeassert
	transport.ResponseHeaderTimeout = ResponseHeaderTimeout

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var target *url.URL

		s := time.Now()
		defer func() {
			d := time.Since(s)
			slog.Debug("relay request completed", "target", target, "duration", d)
			metrics.RelayRequestDuration.Observe(d.Seconds())
		}()

		target, err := getTarget(request)
		if err != nil {
			handleErr(writer, err, http.StatusBadRequest)

			return
		}

		rproxy := &httputil.ReverseProxy{ //nolint:exhaustruct
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

		rproxy.ServeHTTP(writer, request)
	})
}

func getTarget(r *http.Request) (*url.URL, error) {
	t := r.URL.Query().Get(QueryParamTarget)
	if t == "" {
		return nil, fmt.Errorf("%w: %q", ErrQueryParamMissing, QueryParamTarget)
	}

	target, err := url.ParseRequestURI(t)
	if err != nil {
		return nil, fmt.Errorf("target is not a valid URL: %w", err)
	}

	return target, nil
}

func handleErr(w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), status)
	slog.Error("relay handler error", "err", err, "status", status)
	metrics.RelayRequestErrors.Add(1)
}
