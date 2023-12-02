package utils

import (
	"fmt"
	"os"
)

func Failure(message string) {
	fmt.Fprintf(os.Stderr, "failure: %v\n", message)
}
