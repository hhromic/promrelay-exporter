// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package logger_test

import (
	"bytes"
	"context"
	"log/slog"
	"regexp"
	"testing"

	"github.com/hhromic/promrelay-exporter/v2/internal/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandlerString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		handler logger.Handler
		want    string
	}{
		{
			handler: logger.HandlerText,
			want:    "text",
		},
		{
			handler: logger.HandlerJSON,
			want:    "json",
		},
		{
			handler: logger.HandlerTint,
			want:    "tint",
		},
		{
			handler: logger.HandlerAuto,
			want:    "auto",
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.want, tc.handler.String())
	}
}

func TestHandlerMarshalText(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		handler logger.Handler
		want    []byte
		errFunc require.ErrorAssertionFunc
	}{
		{
			handler: logger.HandlerText,
			want:    []byte("text"),
			errFunc: require.NoError,
		},
		{
			handler: logger.HandlerJSON,
			want:    []byte("json"),
			errFunc: require.NoError,
		},
		{
			handler: logger.HandlerTint,
			want:    []byte("tint"),
			errFunc: require.NoError,
		},
		{
			handler: logger.HandlerAuto,
			want:    []byte("auto"),
			errFunc: require.NoError,
		},
	}

	for _, tc := range testCases {
		b, err := tc.handler.MarshalText()
		tc.errFunc(t, err)
		assert.Equal(t, tc.want, b)
	}
}

func TestHandlerUnmarshalText(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		b       []byte
		want    logger.Handler
		errFunc require.ErrorAssertionFunc
	}{
		{
			b:       []byte("text"),
			want:    logger.HandlerText,
			errFunc: require.NoError,
		},
		{
			b:       []byte("json"),
			want:    logger.HandlerJSON,
			errFunc: require.NoError,
		},
		{
			b:       []byte("tint"),
			want:    logger.HandlerTint,
			errFunc: require.NoError,
		},
		{
			b:       []byte("auto"),
			want:    logger.HandlerAuto,
			errFunc: require.NoError,
		},
		{
			b:       []byte("foobar"),
			want:    logger.HandlerText,
			errFunc: require.Error,
		},
	}

	for _, tc := range testCases {
		var h logger.Handler
		err := h.UnmarshalText(tc.b)
		tc.errFunc(t, err)
		assert.Equal(t, tc.want, h)
	}
}

func TestNewSlogLogger(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		handler  logger.Handler
		leveler  slog.Level
		logLevel slog.Level
		logMsg   string
		logArgs  []any
		want     *regexp.Regexp
	}{
		{
			handler:  logger.HandlerText,
			leveler:  slog.LevelDebug,
			logLevel: slog.LevelDebug,
			logMsg:   "message",
			logArgs:  []any{"key1", "val1", "key2", "val2"},
			want:     regexp.MustCompile(`^ts=.+ level=DEBUG msg=message key1=val1 key2=val2\n$`),
		},
		{
			handler:  logger.HandlerJSON,
			leveler:  slog.LevelDebug,
			logLevel: slog.LevelInfo,
			logMsg:   "message",
			logArgs:  []any{"key1", "val1", "key2", "val2"},
			want:     regexp.MustCompile(`^{"ts":".+","level":"INFO","msg":"message","key1":"val1","key2":"val2"}\n$`),
		},
		{
			handler:  logger.HandlerTint,
			leveler:  slog.LevelDebug,
			logLevel: slog.LevelWarn,
			logMsg:   "message",
			logArgs:  []any{"key1", "val1", "key2", "val2"},
			want:     regexp.MustCompile(`^\x1b\[2m.+\x1b\[0m \x1b\[93mWRN\x1b\[0m message \x1b\[2mkey1=\x1b\[0mval1 \x1b\[2mkey2=\x1b\[0mval2\n$`), //nolint:lll
		},
		{
			handler:  logger.HandlerAuto,
			leveler:  slog.LevelDebug,
			logLevel: slog.LevelError,
			logMsg:   "message",
			logArgs:  []any{"key1", "val1", "key2", "val2"},
			want:     regexp.MustCompile(`^ts=.+ level=ERROR msg=message key1=val1 key2=val2\n$`),
		},
		{
			handler:  logger.HandlerText,
			leveler:  slog.LevelWarn,
			logLevel: slog.LevelInfo,
			logMsg:   "message",
			logArgs:  []any{"key1", "val1", "key2", "val2"},
			want:     regexp.MustCompile(`^$`),
		},
	}

	for _, tc := range testCases {
		var buf bytes.Buffer
		l := logger.NewSlogLogger(&buf, tc.handler, tc.leveler)
		require.NotNil(t, l)
		l.Log(context.Background(), tc.logLevel, tc.logMsg, tc.logArgs...)
		assert.Regexp(t, tc.want, buf.String())
	}
}
