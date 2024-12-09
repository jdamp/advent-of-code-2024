package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/jdamp/advent-of-code-2024/util"
	"gonum.org/v1/gonum/mat"
)

//go:embed input.txt
var input string

const NOANTENNA = "."

func parseInput(input string) *util.Grid[string] {
	lines := strings.Split(input, "\n")
	data := make([][]string, len(lines))
	for i, line := range lines {
		data[i] = strings.Split(line, "")
	}
	return util.NewGrid(data)
}

// Find positions of all antennas in the grid
func findAntennas(grid *util.Grid[string]) map[string][]*mat.VecDense {
	antennas := map[string][]*mat.VecDense{}
	for y := 0; y < grid.GetNumRows(); y++ {
		for x := 0; x < grid.GetNumCols(); x++ {
			cell := grid.Get(x, y)
			if cell == NOANTENNA {
				continue
			}
			antennas[cell] = append(antennas[cell], util.AsVector(x, y))
		}
	}
	return antennas
}

func part1(input string) int {
	grid := parseInput(input)
	antennas := findAntennas(grid)
	// For all pairs of antennas (ai, aj) calculate the distance vector aj-ai and find the possible
	// antinodes as ai + 1/2(ai-aj) and aj + 1/2(aj-ai)
	antiNodes := make(map[string]bool)
	for freq := range antennas {
		for i := 0; i < len(antennas[freq]); i++ {
			a := antennas[freq]
			for j := 0; j < i; j++ {
				dist := mat.NewVecDense(2, nil)
				dist.SubVec(a[i], a[j])

				antinode1 := mat.NewVecDense(2, nil)

				antinode2 := mat.NewVecDense(2, nil)
				antinode1.AddScaledVec(a[i], 1, dist)
				antinode2.AddScaledVec(a[j], -1, dist)

				if grid.IsVecValid(antinode1) {
					antiNodes[util.VecToKey(antinode1)] = true
				}

				if grid.IsVecValid(antinode2) {
					antiNodes[util.VecToKey(antinode2)] = true
				}

			}
		}
	}
	return len(antiNodes)
}

func part2(input string) int {
	grid := parseInput(input)
	antennas := findAntennas(grid)
	// Similar to part 1, but now need to find all integer points on the straight line passing
	// through ai and aj. We can do this by finding the gcd of the distance vector and then
	// stepping along the vector in steps of 1/gcd until we hit the edge of the grid.
	antiNodes := make(map[string]bool)
	for freq := range antennas {
		for i := 0; i < len(antennas[freq]); i++ {
			a := antennas[freq]
			for j := 0; j < i; j++ {
				dist := mat.NewVecDense(2, nil)

				dist.SubVec(a[i], a[j])
				gcd := util.GCD(int(dist.AtVec(0)), int(dist.AtVec(1)))
				step := mat.NewVecDense(2, nil)
				step.ScaleVec(1/float64(gcd), dist)

				// 1. Go from ai away from aj until we hit the edge of the grid
				antinode := mat.NewVecDense(2, nil)
				antinode.CloneFromVec(a[i])
				for {
					if !grid.IsVecValid(antinode) {
						break
					}
					antiNodes[util.VecToKey(antinode)] = true
					antinode.AddScaledVec(antinode, -1, step)
				}
				// 2. Go from ai towards aj until we hit the edge of the grid
				antinode.CloneFromVec(a[i])
				for {
					if !grid.IsVecValid(antinode) {
						break
					}
					antiNodes[util.VecToKey(antinode)] = true
					antinode.AddScaledVec(antinode, 1, step)
				}
			}
		}
	}
	return len(antiNodes)
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
