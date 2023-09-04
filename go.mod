module github.com/hhromic/promrelay-exporter/v2

go 1.21

require (
	github.com/alexflint/go-arg v1.4.3
	github.com/go-chi/chi/v5 v5.0.10
	github.com/lmittmann/tint v1.0.1
	github.com/mattn/go-isatty v0.0.19
	github.com/prometheus/client_golang v1.16.0
	go.uber.org/automaxprocs v1.5.3
	golang.org/x/sync v0.3.0
)

require (
	github.com/alexflint/go-scalar v1.2.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/prometheus/client_model v0.4.0 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.11.1 // indirect
	golang.org/x/sys v0.12.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

retract v2.0.0 // Published with wrong module version.
