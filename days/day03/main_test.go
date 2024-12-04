package main

import (
	_ "embed"
	"testing"
)

//go:embed test_input.txt
var test_input string

func TestPart1(t *testing.T) {
	want := 161
	if got := part1(test_input); got != want {
		t.Errorf("part1(): Got %v, want %v", got, want)
	}

}
