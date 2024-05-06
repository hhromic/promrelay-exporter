module github.com/hhromic/promrelay-exporter/v2

go 1.22.0

require (
	github.com/alexflint/go-arg v1.4.3
	github.com/hhromic/go-toolkit v0.0.0-20240325213920-0dc8508dccc1
	github.com/prometheus/client_golang v1.19.0
	go.uber.org/automaxprocs v1.5.3
)

require (
	github.com/alexflint/go-scalar v1.2.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/lmittmann/tint v1.0.4 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.53.0 // indirect
	github.com/prometheus/procfs v0.14.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
)

retract v2.0.0 // Published with wrong module version.
