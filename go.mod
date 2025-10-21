module github.com/bennettzhu1/bazel_go_demo_dep

go 1.23.1

require (
	github.com/bennettzhu1/bazel_go_demo_dep/pkg/api v0.0.0
	github.com/bennettzhu1/bazel_go_demo_dep/pkg/utils v0.0.0-00010101000000-000000000000
)

require github.com/bennettzhu1/go-dep-consumer v0.0.0-20251021051823-63530e754749 // indirect

replace github.com/bennettzhu1/bazel_go_demo_dep/pkg/api => ./pkg/api

replace github.com/bennettzhu1/bazel_go_demo_dep/pkg/utils => ./pkg/utils
