package libs

import (
	"github.com/bennettzhu1/bazel_go_demo/libs/hello"
)

// WrapGreet returns a greeting message
func WrapGreet(audience string) string {
	return hello.Greet(audience)
}
