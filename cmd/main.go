// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/alexflint/go-arg"
	"github.com/hhromic/promrelay-exporter/v2/internal/buildinfo"
	"github.com/hhromic/promrelay-exporter/v2/internal/logger"
	_ "github.com/hhromic/promrelay-exporter/v2/internal/metrics" // initialize collectors
	"github.com/hhromic/promrelay-exporter/v2/internal/server"
	"go.uber.org/automaxprocs/maxprocs"
)

//nolint:lll,tagalign
type args struct {
	ListenAddress string         `arg:"--listen-address,env:LISTEN_ADDRESS" default:":9878" placeholder:"ADDRESS" help:"listen address for the HTTP server"`
	LogHandler    logger.Handler `arg:"--log-handler,env:LOG_HANDLER" default:"auto" placeholder:"HANDLER" help:"application logging handler"`
	LogLevel      slog.Level     `arg:"--log-level,env:LOG_LEVEL" default:"info" placeholder:"LEVEL" help:"application logging level"`
}

func main() {
	var args args

	arg.MustParse(&args)

	if err := logger.SlogSetDefault(os.Stderr, args.LogHandler, args.LogLevel); err != nil {
		panic(err)
	}

	if err := appMain(args); err != nil {
		slog.Error("application error", "err", err)
		os.Exit(1)
	}
}

func appMain(args args) error {
	if _, err := maxprocs.Set(); err != nil {
		slog.Warn("failed to set GOMAXPROCS", "err", err)
	}

	slog.Info("starting",
		"version", buildinfo.Version,
		"goversion", buildinfo.GoVersion,
		"gitcommit", buildinfo.GitCommit,
		"gitbranch", buildinfo.GitBranch,
		"builddate", buildinfo.BuildDate,
		"gomaxprocs", runtime.GOMAXPROCS(0),
	)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	r := server.NewRouter()

	slog.Info("starting HTTP server", "addr", args.ListenAddress)

	if err := server.ListenAndServe(ctx, args.ListenAddress, r); err != nil && !errors.Is(err, context.Canceled) {
		return fmt.Errorf("server listen and serve: %w", err)
	}

	slog.Info("finished")

	return nil
}
