// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	tkhttp "github.com/hhromic/go-toolkit/http"
	"golang.org/x/sync/errgroup"
)

const (
	// ShutdownTimeout is the maximum time to wait for the HTTP server to shutdown.
	ShutdownTimeout time.Duration = 30 * time.Second
	// ReadHeaderTimeout is the maximum time to wait for reading an HTTP request header.
	ReadHeaderTimeout time.Duration = 60 * time.Second
)

// ListenAndServe listens on the TCP network address addr and serves the handler.
// This function implements graceful shutdown when the passed ctx is done.
func ListenAndServe(ctx context.Context, addr string, handler http.Handler) error {
	srv := &http.Server{ //nolint:exhaustruct
		Addr:              addr,
		Handler:           handler,
		ReadHeaderTimeout: ReadHeaderTimeout,
	}

	egrp, ctx := errgroup.WithContext(ctx)

	egrp.Go(func() error {
		return fmt.Errorf("wait and shutdown: %w", tkhttp.WaitAndShutdown(ctx, srv, ShutdownTimeout))
	})

	egrp.Go(func() error {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("listen and serve: %w", err)
		}

		return nil
	})

	if err := egrp.Wait(); err != nil {
		return fmt.Errorf("errgroup wait: %w", err)
	}

	return nil
}

// Error replies to the request with the specified error message and HTTP code.
// It also logs the request remote address, error and code as a warning.
func Error(w http.ResponseWriter, r *http.Request, err string, code int) {
	http.Error(w, err, code)
	slog.Warn("request error", "addr", r.RemoteAddr, "err", err, "code", code)
}
