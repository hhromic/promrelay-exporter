// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"
	"net/http"
	"net/url"
)

const (
	// QueryParamTarget is the request query parameter used for providing a target.
	QueryParamTarget = "target"
)

//nolint:gochecknoglobals
var ctxKeyTargetURL = &contextKey{"targetURL"}

type contextKey struct {
	name string
}

// ExtractQueryParamTargetURL extracts a required target URL from request query parameters.
func ExtractQueryParamTargetURL(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		target := request.URL.Query().Get(QueryParamTarget)
		if target == "" {
			http.Error(writer, "query parameter missing: "+QueryParamTarget, http.StatusBadRequest)

			return
		}

		parsed, err := url.ParseRequestURI(target)
		if err != nil {
			http.Error(writer, QueryParamTarget+" is not a valid URL: "+err.Error(), http.StatusBadRequest)

			return
		}

		request = request.WithContext(context.WithValue(ctx, ctxKeyTargetURL, parsed))

		next.ServeHTTP(writer, request)
	})
}

// TargetURLFromContext returns the target URL value stored in ctx, if any.
func TargetURLFromContext(ctx context.Context) *url.URL {
	if t, ok := ctx.Value(ctxKeyTargetURL).(*url.URL); ok {
		return t
	}

	return nil
}
