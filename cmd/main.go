// SPDX-FileCopyrightText: Copyright 2023 Hugo Hromic
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
	tkslog "github.com/hhromic/go-toolkit/slog"
	"github.com/hhromic/promrelay-exporter/v2/internal/buildinfo"
	_ "github.com/hhromic/promrelay-exporter/v2/internal/metrics" // initialize collectors
	"github.com/hhromic/promrelay-exporter/v2/internal/server"
)

//nolint:lll,tagalign
type args struct {
	ListenAddress string         `arg:"--listen-address,env:LISTEN_ADDRESS" default:":9878" placeholder:"ADDRESS" help:"listen address for the HTTP server"`
	LogHandler    tkslog.Handler `arg:"--log-handler,env:LOG_HANDLER" default:"auto" placeholder:"HANDLER" help:"application logging handler"`
	LogLevel      slog.Level     `arg:"--log-level,env:LOG_LEVEL" default:"info" placeholder:"LEVEL" help:"application logging level"`
}

func (args) Description() string {
	return "Prometheus relay exporter version."
}

func (args) Version() string {
	return buildinfo.Version + " (git:" + buildinfo.GitBranch + "/" + buildinfo.GitCommit + ")"
}

func main() {
	var args args

	arg.MustParse(&args)

	slog.SetDefault(tkslog.NewSlogLogger(os.Stderr, args.LogHandler, args.LogLevel))

	if err := appMain(args); err != nil {
		slog.Error("application error", "err", err)
		os.Exit(1)
	}
}

func appMain(args args) error {
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

	m := server.NewServeMux()

	slog.Info("starting HTTP server", "addr", args.ListenAddress)

	err := server.Run(ctx, args.ListenAddress, m)
	if err != nil && !errors.Is(err, context.Canceled) {
		return fmt.Errorf("run: %w", err)
	}

	slog.Info("finished")

	return nil
}
