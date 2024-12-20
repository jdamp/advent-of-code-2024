package util

import (
	"fmt"
	"strings"

	"gonum.org/v1/gonum/mat"
)

type Grid[T any] struct {
	data [][]T
}

// NewGrid creates a new Grid from a 2D slice
func NewGrid[T any](data [][]T) *Grid[T] {
	return &Grid[T]{data: data}
}

// NewGridFromMultiLineString creates a new Grid from a multi-line string
func NewGridFromMultiLineString(s string) *Grid[string] {
	lines := strings.Split(s, "\n")
	data := make2DSlice[string](len(lines), len(lines[0]))

	for i, line := range lines {
		for j, char := range line {
			data[i][j] = string(char)
		}
	}
	return NewGrid(data)
}

// make2DSlice creates a nRows x mCols slice
func make2DSlice[T any](n int, m int) [][]T {
	newLike := make([][]T, n)
	for i := 0; i < n; i++ {
		newLike[i] = make([]T, m)
	}
	return newLike
}

// NewLike creates a new Grid with the same dimensions as the original Grid filled with a value
func NewConstLike[T, U any](grid *Grid[T], value U) *Grid[U] {
	n := grid.GetNumRows()
	m := grid.GetNumCols()
	data := make2DSlice[U](n, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			data[i][j] = value
		}
	}
	return NewGrid(data)
}

func (grid *Grid[T]) GetNumRows() int {
	return len(grid.data)
}

func (grid *Grid[T]) GetNumCols() int {
	return len(grid.data[0])
}

func (grid *Grid[T]) Get(i, j int) T {
	return grid.data[i][j]
}

func (grid *Grid[T]) Set(i, j int, value T) {
	grid.data[i][j] = value
}

func (grid *Grid[T]) GetData() [][]T {
	return grid.data
}

// Transpose transposes the Grid
func (grid *Grid[T]) Transpose() *Grid[T] {
	n := grid.GetNumRows()
	m := grid.GetNumCols()
	transposed := make2DSlice[T](m, n)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			transposed[j][i] = grid.data[i][j]
		}
	}
	return NewGrid(transposed)
}

// ReverseRows reverses the rows of a Grid
func (grid *Grid[T]) ReverseRows() *Grid[T] {
	n := grid.GetNumRows()
	m := grid.GetNumCols()
	reversed := make2DSlice[T](n, m)

	for i := 0; i < n; i++ {
		reversed[i] = grid.data[n-1-i]
	}
	return NewGrid(reversed)
}

// ReverseCols reverses the columns of a Grid
func (grid *Grid[T]) ReverseCols() *Grid[T] {
	n := grid.GetNumRows()
	m := grid.GetNumCols()
	reversed := make2DSlice[T](n, m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			reversed[i][j] = grid.data[i][m-1-j]
		}
	}
	return NewGrid(reversed)
}

// Rotate90Clockwise rotates the Grid 90 degrees clockwise.
func (grid *Grid[T]) Rotate90Clockwise() *Grid[T] {
	return grid.Transpose().ReverseCols()
}

// IsValid checks if a vector is within the bounds of a Grid
func (g *Grid[T]) IsValid(x, y int) bool {
	return (0 <= x && x < g.GetNumCols() && 0 <= y && y < g.GetNumRows())

}

// IsValid checks if a vector is within the bounds of a Grid
func (g *Grid[T]) IsVecValid(v *mat.VecDense) bool {
	x := int(v.AtVec(0))
	y := int(v.AtVec(1))
	return g.IsValid(x, y)

}

// Print displays the Grid in the terminal
func (g *Grid[T]) Print() {
	for _, row := range g.data {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}
