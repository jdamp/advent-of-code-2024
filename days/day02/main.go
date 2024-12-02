package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/jdamp/advent-of-code-2024/util"
)

//go:embed input.txt
var input string

func parseInput(input string) ([][]int, error) {
	var reports [][]int
	for _, levels := range strings.Split(input, "\n") {
		var report []int
		for _, level := range strings.Fields(levels) {

			value, err := strconv.Atoi(level)
			if err != nil {
				return nil, err
			}
			report = append(report, value)
		}
		reports = append(reports, report)
	}
	return reports, nil
}

func safeDiff(i int) bool {
	abs := util.Abs(i)
	return abs <= 3 && abs > 0
}

func part1(input string) (count int) {
	levels, _ := parseInput(input)
	// Do we have to check whehter the level is decreasing?

	for _, level := range levels {
		diffs, _ := util.IntSliceDiff(level[1:], level[:len(level)-1])
		inc := util.Map(diffs, func(i int) bool { return i > 0 })
		safeDiff := util.Map(diffs, safeDiff)
		check1 := util.All(inc) || !util.Any(inc)
		check2 := util.All(safeDiff)
		if check1 && check2 {
			count += 1
		}

	}
	return count
}

func part2(input string) int {
	// Like part 1, but allow at least one unsafe level
	levels, _ := parseInput(input)
	return levels[0][0]

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
