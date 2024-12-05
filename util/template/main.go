package main

import (
	_ "embed"
	"flag"
	"fmt"
)

//go:embed input.txt
var input string

func parseInput(input string) {

}

func part1(input string) int {
	return 0
}

func part2(input string) int {
	return 0
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
