package util

import (
	"reflect"
	"testing"
)

func NewTestGrid() *Grid[int] {
	return NewGrid([][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	})
}

func TestTranspose(t *testing.T) {
	Grid := NewTestGrid()

	want := [][]int{
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
	}
	transposed := Grid.Transpose()

	if !reflect.DeepEqual(transposed.data, want) {
		t.Errorf("Transpose failed. Wanted %v, got %v", want, transposed)
	}
}

func TestReverseRows(t *testing.T) {
	Grid := NewTestGrid()

	want := [][]int{
		{6, 7, 8},
		{3, 4, 5},
		{0, 1, 2},
	}

	reversed := Grid.ReverseRows()

	if !reflect.DeepEqual(reversed.data, want) {
		t.Errorf("ReverseRows failed. Wanted %v, got %v", want, reversed)
	}
}

func TestReverseCols(t *testing.T) {
	Grid := NewTestGrid()

	want := [][]int{
		{2, 1, 0},
		{5, 4, 3},
		{8, 7, 6},
	}

	reversed := Grid.ReverseCols()

	if !reflect.DeepEqual(reversed.data, want) {
		t.Errorf("ReverseCols failed. Wanted %v, got %v", want, reversed)
	}
}

// Test Rotate90Clockwise method
func TestRotate90Clockwise(t *testing.T) {
	grid := NewTestGrid()

	wanted := [][]int{
		{6, 3, 0},
		{7, 4, 1},
		{8, 5, 2},
	}

	rotated := grid.Rotate90Clockwise()

	if !reflect.DeepEqual(rotated.data, wanted) {
		t.Errorf("Rotate90Clockwise failed. wanted %v, got %v", wanted, rotated.data)
	}
}
