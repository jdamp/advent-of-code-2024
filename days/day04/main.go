package main

import (
	_ "embed"
	"flag"
	"fmt"
	"reflect"
	"strings"

	"github.com/jdamp/advent-of-code-2024/util"
)

//go:embed input.txt
var input string

func ParseInput(input string) *util.Grid[string] {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	gridData := make([][]string, len(lines))

	for i, line := range lines {
		gridData[i] = strings.Split(line, "")
	}

	return util.NewGrid(gridData)
}

// Searches and counts XMAS in a row or a downward right diagonal
func countXmas(grid *util.Grid[string]) (count int) {
	n := grid.GetNumRows()
	m := grid.GetNumCols()
	// After an X, can only see MAS if at least three more values are available
	pattern := []string{"M", "A", "S"}
	for i := 0; i < n; i++ {
		for j := 0; j < m-3; j++ {
			if grid.Get(i, j) != "X" {
				continue
			}
			nextInRow := grid.GetData()[i][j+1 : j+4]
			if reflect.DeepEqual(nextInRow, pattern) {
				count += 1
			}
			// Only check diagonal if sufficient space is available
			if i >= n-3 {
				continue
			}

			nextInDiagonal := []string{
				grid.GetData()[i+1][j+1],
				grid.GetData()[i+2][j+2],
				grid.GetData()[i+3][j+3],
			}
			if reflect.DeepEqual(nextInDiagonal, pattern) {
				count += 1
			}

		}

	}
	return count
}

func part1(input string) (count int) {
	grid := ParseInput(input)
	// Idea: Search for an X, then check for MAS to the left and on the diagonal, then rotate
	count += countXmas(grid)
	for i := 0; i < 3; i++ {
		grid = grid.Rotate90Clockwise()
		count += countXmas(grid)
	}
	return count
}

func part2(input string) (count int) {
	grid := ParseInput(input)
	n := grid.GetNumRows()
	m := grid.GetNumCols()
	// Search for any A's and check the four corners around for S and M in the correct positions
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			if grid.Get(i, j) != "A" {
				continue
			}
			corners := []string{
				grid.Get(i-1, j-1),
				grid.Get(i-1, j+1),
				grid.Get(i+1, j-1),
				grid.Get(i+1, j+1),
			}
			counts := map[string]int{
				"S": 0,
				"M": 0,
			}
			for _, s := range corners {
				if s == "S" || s == "M" {
					counts[s]++
				}
			}
			// Check that there are exactly two "S" and two "M"
			if counts["S"] != 2 || counts["M"] != 2 {
				continue
			}

			// Check that first and last, as well as second and third corners are not the same
			if corners[0] == corners[3] || corners[1] == corners[2] {
				continue
			}
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
