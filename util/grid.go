package util

type Grid[T any] struct {
	data [][]T
}

// NewGrid creates a new Grid from a 2D slice
func NewGrid[T any](data [][]T) *Grid[T] {
	return &Grid[T]{data: data}
}

// make2DSlice creates a nRows x mCols slice
func make2DSlice[T any](n int, m int) [][]T {
	newLike := make([][]T, n)
	for i := 0; i < n; i++ {
		newLike[i] = make([]T, m)
	}
	return newLike
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
