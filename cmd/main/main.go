package main

import (
	"fmt"

	"github.com/bennettzhu1/go-dep-consumer/consumer"
	"github.com/bennettzhu1/bazel_go_demo_dep/pkg/api"
)

func main() {
	fmt.Println("=== Diamond Dependency Test ===")
	fmt.Printf("Direct from pkg/api:    %s\n", api.GetAPIVersion())
	fmt.Printf("Via go-dep-consumer:    %s\n", consumer.GetAPIVersionFromExternal())
	fmt.Println("================================")
}
