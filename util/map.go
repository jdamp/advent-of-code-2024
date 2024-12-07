package util

// Create a copy of a non-nested map
func CopyMap[K comparable, V any](original map[K]V) map[K]V {

	copiedMap := make(map[K]V)

	// Copy key-value pairs from originalMap to copiedMap
	for key, value := range original {
		copiedMap[key] = value
	}
	return copiedMap
}
