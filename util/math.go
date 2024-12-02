package util

import "errors"

// Abs returns the absolute value of an integer.
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// IntSliceDiff calculates the difference between two slices of integers.
func IntSliceDiff(x, y []int) ([]int, error) {
	if len(x) != len(y) {
		return nil, errors.New("slices must have the same length")
	}

	diff := make([]int, len(x))
	for i := range x {
		diff[i] = x[i] - y[i]
	}
	return diff, nil
}
