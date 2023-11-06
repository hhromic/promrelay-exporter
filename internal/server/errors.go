// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package server

import "errors"

// Errors used by the server package.
var (
	// ErrMissingQueryParam is returned when a client request is missing a query parameter.
	ErrMissingQueryParam = errors.New("missing query parameter")
)
