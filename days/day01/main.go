package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/jdamp/advent-of-code-2024/util"
)

//go:embed input.txt
var input string

func parseInput(input string) (left []int, right []int) {
	// Parse the input
	var first, second int
	for _, line := range strings.Split(input, "\n") {
		fmt.Sscanf(line, "%d   %d", &first, &second)

		left = append(left, first)
		right = append(right, second)
	}
	return left, right
}

func part1(input string) (dist int) {
	left, right := parseInput(input)
	// Sort arrays
	sort.Ints(left)
	sort.Ints(right)
	for i := 0; i < len(left); i++ {
		dist += util.Abs(left[i] - right[i])
	}

	return dist

}

func part2(input string) (score int) {
	left, right := parseInput(input)
	counts := make(map[int]int)
	for _, num := range right {
		if count := counts[num]; count > 0 {
			counts[num] += 1
		} else {
			counts[num] = 1
		}
	}
	for _, num := range left {
		if count, ok := counts[num]; ok {
			score += num * count
		}
	}

	return score
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
