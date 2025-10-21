module github.com/bennettzhu1/bazel-go-demo-dep

go 1.23.1

require (
	github.com/bennettzhu1/bazel-go-demo-dep/pkg/api v0.0.0-00010101000000-000000000000
	github.com/bennettzhu1/bazel-go-demo-dep/pkg/utils v0.0.0-00010101000000-000000000000
)

replace github.com/bennettzhu1/bazel-go-demo-dep/pkg/api => ./pkg/api

replace github.com/bennettzhu1/bazel-go-demo-dep/pkg/utils => ./pkg/utils
