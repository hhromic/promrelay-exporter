// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package logger

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"strings"
)

// Handler represents a supported slog handler.
type Handler int

// Names for supported slog handlers.
const (
	// HandlerText represents [slog.TextHandler] which outputs logs in key=value format.
	HandlerText Handler = iota
	// HandlerJSON represents [slog.JSONHandler] which outputs logs in standard JSON format.
	HandlerJSON
)

// Errors used by the logger package.
var (
	// ErrUnknownHandlerName is returned when an unknown slog handler name is used.
	ErrUnknownHandlerName = errors.New("unknown handler name")
)

// String returns a name for the slog handler.
func (h Handler) String() string {
	switch h {
	case HandlerText:
		return "text"
	case HandlerJSON:
		return "json"
	default:
		return ""
	}
}

// MarshalText implements [encoding.TextMarshaler] by calling [Handler.String].
func (h Handler) MarshalText() ([]byte, error) {
	return []byte(h.String()), nil
}

// UnmarshalText implements [encoding.TextUnmarshaler].
// It accepts any string produced by [Handler.MarshalText], ignoring case.
func (h *Handler) UnmarshalText(data []byte) error {
	str := string(data)
	switch strings.ToLower(str) {
	case "text":
		*h = HandlerText
	case "json":
		*h = HandlerJSON
	default:
		return fmt.Errorf("%w: %q", ErrUnknownHandlerName, str)
	}

	return nil
}

// SlogSetDefault sets the global slog logger to output to writer, using the specified log handler
// and the specified minimum logging level. This function also renames the built-in attribute
// [slog.TimeKey] to "ts" for shorter logs output.
func SlogSetDefault(writer io.Writer, handler Handler, level slog.Leveler) error {
	opts := &slog.HandlerOptions{
		AddSource: false,
		Level:     level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Key = "ts"
			}

			return a
		},
	}

	switch handler {
	case HandlerText:
		slog.SetDefault(slog.New(slog.NewTextHandler(writer, opts)))
	case HandlerJSON:
		slog.SetDefault(slog.New(slog.NewJSONHandler(writer, opts)))
	}

	return nil
}
