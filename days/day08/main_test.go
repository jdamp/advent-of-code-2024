package main

import (
	_ "embed"
	"testing"
)

//go:embed test_input.txt
var test_input string

//go:embed test_input2.txt
var test_input2 string

func TestPart1(t *testing.T) {
	want := 14
	got := part1(test_input)
	if got != want {
		t.Errorf("part1(): got %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{test_input, 34},
		{test_input2, 9},
	}
	for _, tt := range testCases {
		got := part2(tt.input)
		if got != tt.want {
			t.Errorf("part2(): got %v, want %v", got, tt.want)
		}
	}
}
