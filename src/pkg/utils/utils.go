package utils

import (
	"fmt"
	"os"
)

func Failure(message string) {
	fmt.Fprintf(os.Stderr, "failure: %v\n", message)
}

func RemoveEmptyStrings(slice []string) []string {
	result := []string{}
	for _, str := range slice {
		if str != "" {
			result = append(result, str)
		}
	}

	return result
}
