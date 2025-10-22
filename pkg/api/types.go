package api

import "github.com/bennettzhu1/bazel_go_demo_dep/pkg/utils"

// GetAPIVersion returns the current API version string
func GetAPIVersion() string {
	return "🚀 LOCAL_MODIFIED 10/21 12:15pm🚀 + " + utils.GetUtilsVersion()
}

