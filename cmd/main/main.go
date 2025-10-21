package main

import (
	"fmt"

	"github.com/bennettzhu1/go-dep-consumer/consumer"
	apiv1 "github.com/bennettzhu1/bazel_go_demo_dep/pkg/api/v1"
)

func main() {
	fmt.Println("=== Diamond Dependency Test ===")
	fmt.Printf("Direct from pkg/api:    %s\n", apiv1.GetAPIVersion())
	fmt.Printf("Via go-dep-consumer:    %s\n", consumer.GetAPIVersionFromExternal())
	fmt.Println("================================")
}
