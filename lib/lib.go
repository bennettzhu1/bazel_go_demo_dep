package lib

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func PrintLibError() {
	fmt.Println("Child Module: Using github.com/pkg/errors v0.9.2 and github.com/google/uuid v1.3.0")
	err := errors.New("an error from child module")
	id := uuid.New()
	fmt.Printf("Error: %v, UUID: %s\n", err, id)
}
