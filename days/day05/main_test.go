package main

import (
	_ "embed"
	"testing"
)

//go:embed test_input.txt
var test_input string

func TestSolve(t *testing.T) {
	want1 := 143
	want2 := 123
	got1, got2 := solve(test_input)
	if got1 != want1 {
		t.Errorf("part1(): got %v, want %v", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("part2(): got %v, want %v", got2, want2)
	}
}
