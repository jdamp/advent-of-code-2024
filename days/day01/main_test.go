package main

import (
	_ "embed"
	"testing"
)

//go:embed test_input.txt
var input1 string

func TestPart1(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Example",
			input: input1,
			want:  11,
		},
	}

	for _, testCase := range testCases {
		if got := part1(testCase.input); got != testCase.want {
			t.Errorf("part1: got %v, want %v", got, testCase.want)
		}
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Example",
			input: input1,
			want:  31,
		},
	}

	for _, testCase := range testCases {
		if got := part2(testCase.input); got != testCase.want {
			t.Errorf("part2: got %v, want %v", got, testCase.want)
		}
	}
}
