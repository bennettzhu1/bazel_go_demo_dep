module github.com/bennettzhu1/multi-module-demo

go 1.23.1

require (
	github.com/pkg/errors v0.9.1
	github.com/bennettzhu1/multi-module-demo/lib v0.0.0-00010101000000-000000000000
)

require github.com/google/uuid v1.3.0 // indirect

replace github.com/bennettzhu1/multi-module-demo/lib => ./lib
