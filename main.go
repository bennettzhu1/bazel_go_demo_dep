package main

import (
	"fmt"
	// _ "crypto/tls"
	"crypto/aes"
	_ "crypto/tls/fipsonly"

	"github.com/bennettzhu1/multi-module-demo/lib"
	"github.com/pkg/errors"
)

func main() {
	fmt.Println("Parent Module: Using github.com/pkg/errors v0.9.1")
	err := errors.New("an error from parent module")
	fmt.Println(err)

	lib.PrintLibError()
	key := []byte("example key 1234")
	_, err = aes.NewCipher(key) // Use = instead of :=
	if err != nil {
		fmt.Println("FIPS compliance check failed:", err)
	} else {
		fmt.Println("FIPS compliance check passed.")
	}

}
