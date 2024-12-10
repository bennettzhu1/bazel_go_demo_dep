package bazel_go_demo_dep

import (
	"github.com/bennettzhu1/bazel_go_demo/libs/hello"
)

// Greet returns a greeting message
func Greet() string {
	return hello.Greet("World")
}
