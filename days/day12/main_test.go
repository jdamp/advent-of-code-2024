package main

import (
	_ "embed"
	"testing"
)

//go:embed test_input1.txt
var test_input1 string

//go:embed test_input2.txt
var test_input2 string

func TestPart1(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{test_input1, 140},
		{test_input2, 1930},
	}
	for _, tt := range testCases {
		got := part1(tt.input)
		if got != tt.want {
			t.Errorf("part1(): got %v, want %v", got, tt.want)
		}
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{test_input1, 80},
		{test_input2, 1206},
	}
	for _, tt := range testCases {
		got := part2(tt.input)
		if got != tt.want {
			t.Errorf("part2(): got %v, want %v", got, tt.want)
		}
	}
}
