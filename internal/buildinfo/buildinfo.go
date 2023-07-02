// Copyright 2023 Hugo Hromic
// SPDX-License-Identifier: Apache-2.0

package buildinfo

import (
	"runtime"
)

// Build info vars populated by GoReleaser.
var (
	BuildDate = "unknown"
	GitBranch = "unknown"
	GitCommit = "unknown"
	Version   = "unknown"
)

var (
	// GoVersion is the version of the Go runtime.
	GoVersion = runtime.Version()
)
