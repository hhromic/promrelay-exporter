// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"log/slog"
	"net/http"
	"net/http/httputil"
	"time"
)

const (
	// ResponseHeaderTimeout is the maximum time to wait for reading an HTTP response header.
	ResponseHeaderTimeout = 60 * time.Second
)

// RelayHandler is an [http.Handler] for target relay requests.
func RelayHandler() http.Handler {
	transport := http.DefaultTransport.(*http.Transport).Clone() //nolint:forcetypeassert
	transport.ResponseHeaderTimeout = ResponseHeaderTimeout

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		target := TargetURLFromContext(ctx)

		rproxy := &httputil.ReverseProxy{ //nolint:exhaustruct
			ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
				slog.Error("relay error", "err", err)
				http.Error(w, "relay error: "+err.Error(), http.StatusBadGateway)
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
