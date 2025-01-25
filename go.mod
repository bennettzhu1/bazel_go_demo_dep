module github.com/bennettzhu1/multi-module-demo

go 1.23.1

require (
	github.com/bennettzhu1/multi-module-demo/lib v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
)

require github.com/google/uuid v1.4.0

replace github.com/bennettzhu1/multi-module-demo/lib => ./lib
