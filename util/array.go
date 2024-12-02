package util

func Map[T, V any](vals []T, fn func(T) V) []V {
	result := make([]V, len(vals))
	for i, val := range vals {
		result[i] = fn(val)
	}
	return result
}

func All(vals []bool) bool {
	for _, val := range vals {
		if !val {
			return false
		}
	}
	return true
}

func Any(vals []bool) bool {
	for _, val := range vals {
		if val {
			return true
		}
	}
	return false
}
