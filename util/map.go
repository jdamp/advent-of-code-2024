package util

import "fmt"

// Create a copy of a non-nested map
func CopyMap[K comparable, V any](original map[K]V) map[K]V {

	copiedMap := make(map[K]V)

	// Copy key-value pairs from originalMap to copiedMap
	for key, value := range original {
		copiedMap[key] = value
	}
	return copiedMap
}

// Formats a slice as a string to be used as a key in a map
func SliceAsKey(vals []int) string {
	result := ""
	for i, val := range vals {
		if i > 0 {
			result += ";"
		}
		result += fmt.Sprintf("%d", val)
	}
	return result
}
