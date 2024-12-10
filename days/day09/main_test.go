package main

import (
	_ "embed"
	"testing"
)

//go:embed test_input.txt
var test_input string

//go:embed test_input2.txt
var test_input2 string

func TestParseInput(t *testing.T) {
	want := "00...111...2...333.44.5555.6666.777.888899"
	gotMap := parseInput(test_input)
	for i, c := range want {
		if gotMap[i] != string(c) {
			t.Errorf("parseInput(): got %v, want %v", gotMap[i], string(c))
		}
	}
}

func TestMoveBlocks(t *testing.T) {
	want := "0099811188827773336446555566.............."
	gotMap := parseInput(test_input)
	gotMap = moveBlocks(gotMap)
	for i, c := range want {
		if gotMap[i] != string(c) {
			t.Errorf("moveBlocks(): got %v, want %v", gotMap[i], string(c))
		}
	}
}

func TestMoveFiles(t *testing.T) {
	want := "00992111777.44.333....5555.6666.....8888.."
	gotMap := parseInput(test_input)
	moved := moveFiles(gotMap)
	for i, c := range want {
		if moved[i] != string(c) {

			t.Errorf("moveFiles(): got %v, want %v", moved[i], string(c))
		}
	}
}

func TestPart1(t *testing.T) {
	want := 1928
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
		{test_input, 2858},
		{"121", 1},
		{"14113", 16},
	}
	for _, tt := range testCases {
		got := part2(tt.input)
		if got != tt.want {
			t.Errorf("part2(): got %v, want %v", got, tt.want)
		}
	}
}
