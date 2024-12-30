package main

import (
	"fmt"

	"github.com/bennettzhu1/multi-module-demo/lib"
	"github.com/pkg/errors"
)

func main() {
	fmt.Println("Parent Module: Using github.com/pkg/errors v0.9.1")
	err := errors.New("an error from parent module")
	fmt.Println(err)

	lib.PrintLibError()
}
