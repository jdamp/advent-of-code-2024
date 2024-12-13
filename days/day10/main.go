package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/jdamp/advent-of-code-2024/util"
)

//go:embed input.txt
var input string

func parseInput(input string) (*util.Grid[int], [][]int) {
	var data [][]int       // store the grid
	var trailheads [][]int // store the trailheads
	for i, line := range strings.Split(input, "\n") {
		data = append(data, []int{})
		for j, char := range line {
			height := int(char - '0')
			data[i] = append(data[i], height)
			if height == 0 {
				trailheads = append(trailheads, []int{i, j})
			}
		}
	}
	return util.NewGrid(data), trailheads
}

// Runs a breadth-first search over the grid starting at x0, yo
func BFS(grid *util.Grid[int], x0, y0 int) int {
	// Initialize the queue with the starting point
	queue := [][]int{{x0, y0}}
	tops := map[string]bool{}
	deltas := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	// Run the BFS
	// Check all four neighbors of the current trailhead
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		curr := grid.Get(head[0], head[1])
		for _, d := range deltas {
			dx, dy := d[0], d[1]
			x, y := head[0]+dx, head[1]+dy
			if !grid.IsValid(x, y) {
				continue
			}
			next := grid.Get(x, y)
			if next == curr+1 && next == 9 {
				tops[util.SliceAsKey([]int{x, y})] = true
			} else if next == curr+1 {
				queue = append(queue, []int{x, y})
			}
		}
	}
	return len(tops)
}

func part1(input string) (score int) {
	grid, trailheads := parseInput(input)
	// Run a breadth-first search over the trailheads
	for _, trailhead := range trailheads {
		score += BFS(grid, trailhead[0], trailhead[1])
	}
	return score
}

// Runs a depth-first search over the grid starting at x, y
func DFS(grid *util.Grid[int], x, y, target int) (paths int) {
	// Recursive approach
	// Check if the current node is valid
	if !grid.IsValid(x, y) || grid.Get(x, y) != target {
		return 0
	}
	// Reached the end of a path
	if target == 9 {
		return 1
	}

	deltas := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range deltas {
		dx, dy := d[0], d[1]
		paths += DFS(grid, x+dx, y+dy, target+1)
	}
	return paths
}

func part2(input string) (score int) {
	grid, trailheads := parseInput(input)
	// Run a breadth-first search over the trailheads
	for _, trailhead := range trailheads {
		score += DFS(grid, trailhead[0], trailhead[1], 0)
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
