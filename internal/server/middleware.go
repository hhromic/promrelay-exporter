// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/hhromic/promrelay-exporter/internal/metrics"
	"golang.org/x/exp/slog"
)

type contextKey string

var (
	targetCtxKey = contextKey("Target")
)

// RelayTargetExtractor is a middleware for extracting a 'target' query parameter from the request.
// The extracted value (if any) is parsed as a *url.URL object. If the parameter is missing or
// invalid, the middleware returns a bad request error to the client.
func RelayTargetExtractor(next http.Handler) http.Handler {
	handleErr := func(w http.ResponseWriter, err error, status int) {
		http.Error(w, err.Error(), status)
		slog.Error("target extractor error", "err", err, "status", status)
		metrics.RelayRequestErrors.Add(1)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		if len(query["target"]) != 1 || query["target"][0] == "" {
			err := fmt.Errorf("'target' parameter is missing or is specified multiple times")
			handleErr(w, err, http.StatusBadRequest)
			return
		}

		target, err := url.ParseRequestURI(query["target"][0])
		if err != nil {
			err := fmt.Errorf("'target' parameter is not a valid URL: %w", err)
			handleErr(w, err, http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		ctx = NewTargetContext(ctx, target)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// NewTargetContext returns a new Context that carries value t.
func NewTargetContext(ctx context.Context, t *url.URL) context.Context {
	return context.WithValue(ctx, targetCtxKey, t)
}

// TargetFromContext returns the target value stored in ctx, if any.
func TargetFromContext(ctx context.Context) (*url.URL, bool) {
	t, ok := ctx.Value(targetCtxKey).(*url.URL)
	return t, ok
}

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
