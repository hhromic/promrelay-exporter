// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

const (
	// QueryParamTarget is the request query parameter used for providing the relay target.
	QueryParamTarget = "target"
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
		target, err := getTarget(request)
		if err != nil {
			Error(writer, request, err.Error(), http.StatusBadRequest)

			return
		}

		rproxy := &httputil.ReverseProxy{ //nolint:exhaustruct
			ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
				Error(w, r, "relay error: "+err.Error(), http.StatusBadGateway)
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
		return nil, fmt.Errorf("%w: %q", ErrMissingQueryParam, QueryParamTarget)
	}

	target, err := url.ParseRequestURI(t)
	if err != nil {
		return nil, fmt.Errorf("target is not a valid URL: %w", err)
	}

	return target, nil
}
