module github.com/bennettzhu1/bazel_go_demo_dep

go 1.23.1

require github.com/bennettzhu1/bazel_go_demo_dep/pkg/api v0.0.0

require github.com/bennettzhu1/go-dep-consumer v0.0.0-20251021061510-460c05c39440

replace github.com/bennettzhu1/bazel_go_demo_dep/pkg/api => ./pkg/api

replace github.com/bennettzhu1/bazel_go_demo_dep/pkg/utils => ./pkg/utils
