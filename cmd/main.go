// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/alexflint/go-arg"
	"github.com/hhromic/promrelay-exporter/internal/buildinfo"
	"github.com/hhromic/promrelay-exporter/internal/logger"
	"github.com/hhromic/promrelay-exporter/internal/server"
	"go.uber.org/automaxprocs/maxprocs"
	"golang.org/x/exp/slog"

	_ "github.com/hhromic/promrelay-exporter/internal/metrics" // initialize collectors
)

type args struct {
	ListenAddress string         `arg:"--listen-address,env:LISTEN_ADDRESS" default:":9878" placeholder:"ADDRESS" help:"listen address for the HTTP server"`
	LogHandler    logger.Handler `arg:"--log-handler,env:LOG_HANDLER" default:"text" placeholder:"HANDLER" help:"application logging handler"`
	LogLevel      slog.Level     `arg:"--log-level,env:LOG_LEVEL" default:"info" placeholder:"LEVEL" help:"application logging level"`
}

func main() {
	var args args
	arg.MustParse(&args)

	if err := logger.SlogSetDefault(os.Stderr, args.LogHandler, args.LogLevel); err != nil {
		panic(err)
	}

	_, err := maxprocs.Set()
	if err != nil {
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
		slog.Error("error running HTTP server", "err", err)
	} else {
		slog.Info("finished")
	}
}
