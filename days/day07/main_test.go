package main

import (
	_ "embed"
	"testing"
)

//go:embed test_input.txt
var test_input string

func TestPart1(t *testing.T) {
	var want int64 = 3749
	got := part1(test_input)
	if got != want {
		t.Errorf("part1(): got %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	var want int64 = 11387
	got := part2(test_input)
	if got != want {
		t.Errorf("part2(): got %v, want %v", got, want)
	}
}
