package main

import (
	_ "embed"
	"flag"
	"fmt"

	"github.com/jdamp/advent-of-code-2024/util"
)

//go:embed input.txt
var input string

var directions = [][2]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func findPatches(grid *util.Grid[string]) [][][2]int {
	visited := util.NewConstLike(grid, false)
	var patches [][][2]int
	// Iterate over the grid
	for y := 0; y < grid.GetNumRows(); y++ {
		for x := 0; x < grid.GetNumCols(); x++ {
			if visited.Get(x, y) {
				continue
			}
			patch := bfs(grid, visited, x, y)
			patches = append(patches, patch)
		}
	}
	return patches
}

// Calculates the perimeter of a patch
func getPerimeter(grid *util.Grid[string], patch [][2]int) int {
	perimeter := 0
	for _, pos := range patch {
		x, y := pos[0], pos[1]
		edges := 4
		for _, dir := range directions {
			newX, newY := x+dir[0], y+dir[1]
			if grid.IsValid(newX, newY) && grid.Get(newX, newY) == grid.Get(x, y) {
				edges--
			}
		}
		perimeter += edges
	}
	return perimeter
}

// Idea for part 2: We need to find corners, as a side is always delimited by corners
// There are two possible kind of corners:
// inward /concave corners
// outward/ convext corners
// These all are identified by the values of their neighboring cells, correctly taking out of
// bound cells intoaccount
func countSides[T comparable](grid *util.Grid[T], patch [][2]int) int {
	numCorners := 0
	directions := map[string][2]int{
		"left":        {0, -1},
		"left above":  {-1, -1},
		"above":       {-1, 0},
		"right above": {-1, 1},
		"right":       {0, 1},
		"right below": {1, 1},
		"below":       {1, 0},
		"left below":  {1, -1},
	}
	for _, pos := range patch {
		// Need to check all eight surrounding directions whether they have the same value as
		// the current position
		x, y := pos[0], pos[1]
		hasSame := make(map[string]bool)
		for key := range directions {
			hasSame[key] = false
		}

		value := grid.Get(x, y)
		for name, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			if !grid.IsValid(nx, ny) {
				continue
			}
			if value == grid.Get(nx, ny) {
				hasSame[name] = true
			}
		}
		// Use the hasSame map to check all possible corner types
		// Upper left convex
		if !hasSame["above"] && !hasSame["right"] {
			numCorners++
		}
		// Upper right convex
		if !hasSame["above"] && !hasSame["left"] {
			numCorners++
		}
		// Lower right convex
		if !hasSame["below"] && !hasSame["right"] {
			numCorners++
		}
		// Lower left convex
		if !hasSame["below"] && !hasSame["left"] {
			numCorners++
		}
		// Upper left concave
		if hasSame["above"] && hasSame["left"] && !hasSame["left above"] {
			numCorners++
		}
		// Upper right concave
		if hasSame["above"] && hasSame["right"] && !hasSame["right above"] {
			numCorners++
		}
		// Lower right concave
		if hasSame["below"] && hasSame["right"] && !hasSame["right below"] {
			numCorners++
		}
		// Lower left concave
		if hasSame["below"] && hasSame["left"] && !hasSame["left below"] {
			numCorners++
		}
	}
	return numCorners
}

// Finds all the points belonging to a region
func bfs(grid *util.Grid[string], visited *util.Grid[bool], startX, startY int) [][2]int {
	queue := [][2]int{{startX, startY}}
	region := [][2]int{}
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		x, y := pos[0], pos[1]
		if visited.Get(x, y) {
			continue
		}
		// Mark the cell as visited and add it to the patchj
		visited.Set(x, y, true)
		region = append(region, [2]int{x, y})
		for _, dir := range directions {
			newX, newY := pos[0]+dir[0], pos[1]+dir[1]

			if grid.IsValid(newX, newY) && !visited.Get(newX, newY) && grid.Get(newX, newY) == grid.Get(x, y) {
				queue = append(queue, [2]int{newX, newY})
			}

		}
	}
	return region
}

func part1(input string) (price int) {
	grid := util.NewGridFromMultiLineString(input)
	// Use a breadth-first search to find the patches
	patches := findPatches(grid)
	for _, patch := range patches {
		perimeter := getPerimeter(grid, patch)
		price += perimeter * len(patch)
	}
	return price
}

func part2(input string) (price int) {
	grid := util.NewGridFromMultiLineString(input)
	// Use a breadth-first search to find the patches
	patches := findPatches(grid)
	for _, patch := range patches {
		sides := countSides(grid, patch)
		price += sides * len(patch)
	}
	return price
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
