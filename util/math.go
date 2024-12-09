package util

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

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

// GCD returns the greatest common divisor of two integers.
func GCD(a, b int) int {
	if b == 0 {
		return a
	} else {
		return GCD(b, a%b)
	}
}

// VecToKey converts a vector into a string key for storage in a map
func VecToKey(vec *mat.VecDense) string {
	var builder strings.Builder
	for _, val := range vec.RawVector().Data {
		builder.WriteString(fmt.Sprintf("%d,", int(val)))
	}
	return builder.String()
}

// KeyToVecx converts a string created by VecToKey back into a vector
func KeyToVec(key string) *mat.VecDense {
	parts := strings.Split(key, ",")
	data := make([]float64, 2)
	for i, part := range parts {
		val, _ := strconv.Atoi(part)
		data[i] = float64(val)
	}
	return mat.NewVecDense(len(data), data)
}

// AsVector creates a new vector from two integers
func AsVector(x, y int) *mat.VecDense {
	return mat.NewVecDense(2, []float64{float64(x), float64(y)})
}
