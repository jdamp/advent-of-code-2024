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

func validateLevel(level []int) bool {
	diffs, _ := util.IntSliceDiff(level[1:], level[:len(level)-1])
	isInc := util.Map(diffs, func(i int) bool { return 3 >= i && i > 0 })
	isDec := util.Map(diffs, func(i int) bool { return -3 <= i && i < 0 })
	return util.All(isInc) || util.All(isDec)
}

func validateWithRemoval(level []int) bool {
	for i := 0; i < len(level); i++ {
		levelWithOut := append([]int{}, level[:i]...)
		levelWithOut = append(levelWithOut, level[i+1:]...)
		if validateLevel(levelWithOut) {
			return true
		}
	}
	return false
}

func part1(input string) (count int) {
	levels, _ := parseInput(input)
	for _, level := range levels {
		if validateLevel(level) {
			count += 1
		}
	}
	return count
}

func part2(input string) (count int) {
	levels, _ := parseInput(input)
	for _, level := range levels {
		// Like part 1, but allow at least one unsafe level
		// 1000 input lines with only a few numbers per line, so the brute-force approach
		// should be sufficient
		if validateLevel(level) || validateWithRemoval(level) {
			count += 1
		}
	}
	return count
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
