package libs

import (
	"github.com/bennettzhu1/bazel_go_demo/libs/hello"
)

// WrapGreet returns a greeting message
func WrapGreet() string {
	return hello.Greet("World")
}
