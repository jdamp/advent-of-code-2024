package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

func part1(input string) (result int) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		first, _ := strconv.Atoi(match[1])
		second, _ := strconv.Atoi(match[2])
		result += first * second

	}
	return result
}

func part2(input string) (result int) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	mul := 1 // 1 if enabled 0 if disabled
	for _, match := range matches {
		if match[0] == "do()" {
			mul = 1
			continue
		} else if match[0] == "don't()" {
			mul = 0
			continue
		}
		first, _ := strconv.Atoi(match[1])
		second, _ := strconv.Atoi(match[2])
		result += mul * first * second
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
