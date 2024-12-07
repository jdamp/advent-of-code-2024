package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/jdamp/advent-of-code-2024/util"
	"gonum.org/v1/gonum/mat"
)

//go:embed input.txt
var input string

type State struct {
	position  *mat.VecDense
	velocity  *mat.VecDense
	xmax      int
	ymax      int
	obstacles map[string]bool
}

// Since go doesn't have a set type, we'll use a map to store the unique visited coordinates
// again. vecToKey converts a vector into a string key for the visited and obstacles map
func vecToKey(vec *mat.VecDense) string {
	var builder strings.Builder
	for _, val := range vec.RawVector().Data {
		builder.WriteString(fmt.Sprintf("%d,", int(val)))
	}
	return builder.String()
}

func keyToVec(key string) *mat.VecDense {
	parts := strings.Split(key, ",")
	data := make([]float64, 2)
	for i, part := range parts {
		val, _ := strconv.Atoi(part)
		data[i] = float64(val)
	}
	return mat.NewVecDense(len(data), data)
}

func (s *State) inBound() bool {
	x := s.position.AtVec(0)
	y := s.position.AtVec(1)
	return (0 <= x && x <= float64(s.xmax)) && (0 <= y && y <= float64(s.ymax))
}

func asVector(x, y int) *mat.VecDense {
	return mat.NewVecDense(2, []float64{float64(x), float64(y)})
}

func parseInput(input string) State {
	state := State{
		obstacles: make(map[string]bool),
	}
	rows := strings.Split(input, "\n")
	state.xmax = len(rows[0]) - 1
	state.ymax = len(rows) - 1
	for x := 0; x <= state.xmax; x++ {
		for y := 0; y <= state.ymax; y++ {
			sym := rows[y][x]
			if sym == '^' {
				state.position = asVector(x, y)
				state.velocity = asVector(0, -1)
			} else if sym == '#' {
				state.obstacles[vecToKey(asVector(x, y))] = true
			}
		}
	}
	return state
}

func rotate(velocity *mat.VecDense) *mat.VecDense {
	newVelocity := mat.NewVecDense(2, nil)
	m := mat.NewDense(2, 2, []float64{0, -1, 1, 0})
	newVelocity.MulVec(m, velocity)
	return newVelocity
}

func simulatePart1(state *State) map[string]bool {
	visited := map[string]bool{}
	for state.inBound() {
		visited[vecToKey(state.position)] = true
		simulateStep(state, state.obstacles)
	}
	return visited
}

func simulatePart2(state *State, originalPath map[string]bool) int {
	loopCount := 0

	// Try placing an obstruction at each position in the path
	for pos := range originalPath {
		// Skip the starting position
		if pos == vecToKey(state.position) {
			continue
		}

		// Place the new obstacle and reset the simulation
		obstacles := util.CopyMap(state.obstacles)
		obstacles[pos] = true

		if causesLoop(state, obstacles) {
			loopCount++
		}
	}

	return loopCount
}

func causesLoop(state *State, obstacles map[string]bool) bool {
	visited := map[string]string{}
	simulationState := *state

	for simulationState.inBound() {
		positionKey := vecToKey(simulationState.position)
		velocityKey := vecToKey(simulationState.velocity)

		// Detect loop: the same position and velocity revisited
		if visited[positionKey] == velocityKey {
			return true
		}

		// Mark the position and velocity as visited
		visited[positionKey] = velocityKey

		// Simulate one step
		if !simulateStep(&simulationState, obstacles) {
			break // Guard moved out of bounds or got stuck
		}
	}

	return false
}

func simulateStep(state *State, obstacles map[string]bool) bool {
	newPosition := mat.NewVecDense(2, nil)
	initialVelocity := vecToKey(state.velocity)

	for {
		newPosition.AddVec(state.position, state.velocity)

		if !isObstacle(newPosition, obstacles) {
			state.position = newPosition
			return true // Successfully moved
		}
		state.velocity = rotate(state.velocity)

		// If velocity cycles back to the original without moving, the guard is stuck
		if vecToKey(state.velocity) == initialVelocity {
			return false
		}
	}
}

func isObstacle(position *mat.VecDense, obstacles map[string]bool) bool {
	_, exists := obstacles[vecToKey(position)]
	return exists
}

func part1(input string) map[string]bool {
	state := parseInput(input)
	return simulatePart1(&state)
}

func part2(input string, originalPath map[string]bool) int {
	state := parseInput(input)
	return simulatePart2(&state, originalPath)
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "run which part (1 or 2)")
	flag.Parse()

	if part == 1 {
		visited := part1(input)
		fmt.Println("Result: ", len(visited))
	} else {
		visited := part1(input) // Part 2 relies on the path from part 1
		result := part2(input, visited)
		fmt.Println("Result: ", result)
	}
}
