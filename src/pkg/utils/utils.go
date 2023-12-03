package utils

import (
	"fmt"
	"os"
)

func Failure(message string) {
	fmt.Fprintf(os.Stderr, "failure: %v\n", message)
}

func IsDigit(c byte) bool {
	return c >= 48 && c <= 57
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
