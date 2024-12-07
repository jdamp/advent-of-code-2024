package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/jdamp/advent-of-code-2024/util"
)

//go:embed input.txt
var input string

func solve(input string) (part1, part2 int) {
	// Separate rules and pages
	parts := strings.Split(input, "\n\n")
	rules := strings.Split(parts[0], "\n")
	allPages := strings.Split(parts[1], "\n")

	ruleBook := buildRuleBook(rules)

	for _, pages := range allPages {
		pageSequence := strings.Split(pages, ",")

		// Part1: Calculate score of isValid sequences
		score, isValid := calculateScore(ruleBook, pageSequence)
		part1 += score

		// Part 2: Handle invalid sequences
		if !isValid {
			graph := buildGraph(pageSequence, rules)
			topoOrdered, _ := graph.TopologicalSort()

			// Add the middle value as score for part2
			middle, _ := strconv.Atoi(topoOrdered[len(topoOrdered)/2])
			part2 += middle
		}
	}
	return part1, part2
}

// Create a rulebook from the list of rules
func buildRuleBook(rules []string) map[string][]string {
	ruleBook := make(map[string][]string)
	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		ruleBook[parts[0]] = append(ruleBook[parts[0]], parts[1])
	}
	return ruleBook
}

// Checks if a page sequence is valid according to the ruleBook and calculates the score
func calculateScore(ruleBook map[string][]string, sequence []string) (score int, isValid bool) {
	isValid = true

	for i := 0; i < len(sequence); i++ {
		for j := 0; j < i; j++ {
			if !slices.Contains(ruleBook[sequence[j]], sequence[i]) {
				isValid = false
				break
			}
		}
	}
	if isValid {
		score, _ = strconv.Atoi(sequence[len(sequence)/2])
	}

	return score, isValid
}

// buildGraph parses a list of page numbers and the corresponding rules into a Graph
func buildGraph(sequence []string, rules []string) *util.Graph[string] {
	graph := util.NewGraph[string]()

	for _, rule := range rules {
		parts := strings.Split(rule, "|")

		// Only add edge to graph if both nodes are part of the current sequence

		if slices.Contains(sequence, parts[0]) && slices.Contains(sequence, parts[1]) {
			graph.AddEdge(parts[0], parts[1])
		}
	}
	return graph
}

func main() {
	part1, part2 := solve(input)
	fmt.Println("Result: ", part1)
	fmt.Println("Result: ", part2)

}
