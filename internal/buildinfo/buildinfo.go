// SPDX-FileCopyrightText: Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package buildinfo

import (
	"runtime"
)

// Build info vars populated by GoReleaser.
//
//nolint:gochecknoglobals
var (
	BuildDate = "unknown"
	GitBranch = "unknown"
	GitCommit = "unknown"
	Version   = "unknown"
)

// GoVersion is the version of the Go runtime.
//
//nolint:gochecknoglobals
var GoVersion = runtime.Version()
