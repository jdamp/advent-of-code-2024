package main

import (
	_ "embed"
	"testing"
)

//go:embed test_input_part1.txt
var test_input_part1 string

//go:embed test_input_part2.txt
var test_input_part2 string

func TestPart1(t *testing.T) {
	want := 161
	if got := part1(test_input_part1); got != want {
		t.Errorf("part1(): Got %v, want %v", got, want)
	}

}

func TestPart2(t *testing.T) {
	want := 48
	if got := part2(test_input_part2); got != want {
		t.Errorf("part1(): Got %v, want %v", got, want)
	}

}
