module github.com/hhromic/promrelay-exporter/v2

go 1.21

require (
	github.com/alexflint/go-arg v1.4.3
	github.com/go-chi/chi/v5 v5.0.10
	github.com/hhromic/go-toolkit v0.0.0-20231024233041-414f47bca371
	github.com/prometheus/client_golang v1.17.0
	go.uber.org/automaxprocs v1.5.3
	golang.org/x/sync v0.4.0
)

require (
	github.com/alexflint/go-scalar v1.2.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/lmittmann/tint v1.0.2 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.45.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/term v0.13.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

retract v2.0.0 // Published with wrong module version.
