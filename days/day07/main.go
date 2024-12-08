package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

// Define a type for the operator functions
type operator func(int64, int64) int64

// Define a type to hold all equation-related data
type Equation struct {
	result int64
	values []int64
}

var add = func(a, b int64) int64 { return a + b }
var mult = func(a, b int64) int64 { return a * b }
var concat = func(a, b int64) int64 {
	strConcat := fmt.Sprintf("%d", a) + fmt.Sprintf("%d", b)
	intConcat, _ := strconv.ParseInt(strConcat, 10, 64)
	return intConcat
}

// Parse each line in the input into a Equation
func parseInput(input string, ops []operator) (equations []Equation) {

	for _, line := range strings.Split(input, "\n") {
		equations = append(equations, *NewEquation(line))
	}
	return equations
}

// NewEquation creates a new Equation from a single input line.
func NewEquation(line string) *Equation {
	parts := strings.Split(line, ": ")
	result, _ := strconv.ParseInt(parts[0], 10, 64)

	values := []int64{}
	valuesStr := strings.Fields(parts[1])
	for _, value := range valuesStr {
		v, _ := strconv.ParseInt(value, 10, 64)
		values = append(values, v)
	}
	return &Equation{result: result, values: values}
}

func isSolvable(result int64, candidate int64, values []int64, ops []operator) bool {
	if candidate > result {
		return false
	}
	if len(values) == 0 {
		return candidate == result
	}
	for _, op := range ops {
		if isSolvable(result, op(candidate, values[0]), values[1:], ops) {
			return true
		}
	}
	return false
}

func part1(input string) (result int64) {
	operators := []operator{add, mult}
	equations := parseInput(input, operators)
	for _, eq := range equations {
		if isSolvable(eq.result, eq.values[0], eq.values[1:], operators) {
			result += eq.result
		}
	}
	return result
}

func part2(input string) (result int64) {
	operators := []operator{add, mult, concat}
	equations := parseInput(input, operators)
	for _, eq := range equations {
		if isSolvable(eq.result, eq.values[0], eq.values[1:], operators) {
			result += eq.result
		}
	}
	return result
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "run which part (1 or 2)")
	flag.Parse()
	if part == 1 {
		result := part1(input)
		fmt.Println("Result: ", result)
	} else {
		result := part2(input)
		fmt.Println("Result: ", result)
	}

}
