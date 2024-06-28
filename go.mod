module github.com/hhromic/promrelay-exporter/v2

go 1.22.0

require (
	github.com/alexflint/go-arg v1.5.1
	github.com/hhromic/go-toolkit v0.0.0-20240528221625-da900e00522e
	github.com/prometheus/client_golang v1.19.1
	go.uber.org/automaxprocs v1.5.3
)

require (
	github.com/alexflint/go-scalar v1.2.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/lmittmann/tint v1.0.4 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.55.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	golang.org/x/sys v0.21.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)

retract v2.0.0 // Published with wrong module version.
