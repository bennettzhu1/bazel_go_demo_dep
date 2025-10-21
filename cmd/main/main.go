package main

import (
	"fmt"

	apiv1 "github.com/bennettzhu1/bazel-go-demo-dep/pkg/api/v1"
	utilsv1 "github.com/bennettzhu1/bazel-go-demo-dep/pkg/utils/v1"
)

func main() {
	fmt.Println("=== Bazel Go Subpackage Test ===")
	fmt.Printf("API Version: %s\n", apiv1.GetAPIVersion())
	fmt.Printf("Utils Version: %s\n", utilsv1.GetUtilsVersion())
	fmt.Println("=================================")
}
