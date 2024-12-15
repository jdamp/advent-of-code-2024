package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jdamp/advent-of-code-2024/util"
)

//go:embed input.txt
var input string

func parseInput(input string) map[int]int {
	data := map[int]int{}
	for _, str := range strings.Split(input, " ") {
		// Do something with line
		num, _ := strconv.Atoi(str)
		util.InsertOrIncrementByValue(data, num, 1)
	}
	return data
}

func blink(state map[int]int) map[int]int {
	var newState = map[int]int{}
	var result int
	second := -1
	for num, count := range state {
		if num == 0 {
			result = 1
		} else if len := ndigits(num); len%2 == 0 {
			// x /10^n = a right shift by n digits => divide by 10^(len/2)
			result = num / int(math.Pow10(len/2))
			second = num % int(math.Pow10(len/2))

		} else {
			result = num * 2024
		}
		util.InsertOrIncrementByValue(newState, result, count)
		if second != -1 {
			util.InsertOrIncrementByValue(newState, second, count)
			second = -1
		}
	}
	return newState
}

func ndigits(num int) int {
	return int(math.Floor(math.Log10(float64(num)))) + 1
}

func run(input string, n int) int {
	state := parseInput(input)
	for i := 0; i < n; i++ {
		state = blink(state)
	}
	result := 0
	// Sum all values in state
	for _, count := range state {
		result += count
	}
	return result
}

func part1(input string) int {
	return run(input, 25)
}

func part2(input string) int {
	return run(input, 75)
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
