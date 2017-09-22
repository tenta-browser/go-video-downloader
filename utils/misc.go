package utils

import "fmt"

// Debug mode permits logging
const Debug = false

// Log logs a message from a running extractor
func Log(msg string, args ...interface{}) {
	if Debug {
		fmt.Printf(msg+"\n", args...)
	}
}

// MakeSet creates a map from the specified keys;
// the values are irrelevant, this just wants to simulate a set
func MakeSet(values []string) map[string]struct{} {
	res := make(map[string]struct{})
	for _, value := range values {
		res[value] = struct{}{}
	}
	return res
}
