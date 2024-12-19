package matrix

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/constraints"
	"io"
	"iter"
)

type Matrix2D[T constraints.Integer | constraints.Float] struct {
	data    [][]T
	Rows    int
	Columns int
}

func NewMatrix2D[T constraints.Integer | constraints.Float](data [][]T) Matrix2D[T] {
	X := -1
	for x, b := range data {
		if x == 0 {
			X = len(b)
		} else if X != len(b) {
			panic("Data is not a square 2D array")
		}
	}
	return Matrix2D[T]{data: data, Rows: len(data), Columns: X}
}

func NewEmptyMatrix2D[T constraints.Integer | constraints.Float](rows, columns int) Matrix2D[T] {
	var r [][]T
	for range rows {
		row := make([]T, columns)
		r = append(r, row)
	}
	return Matrix2D[T]{data: r, Rows: rows, Columns: columns}
}

func ReadMatrix2D(reader io.Reader) Matrix2D[rune] {
	scanner := bufio.NewScanner(reader)
	var lines [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, []rune(line))
	}
	return NewMatrix2D(lines)
}

func (t Matrix2D[T]) IterRows() iter.Seq2[int, Vector[T]] {
	return func(yield func(int, Vector[T]) bool) {
		for row := 0; row < t.Rows; row++ {
			arr := t.data[row]
			if !yield(row, Vector[T]{data: arr, size: t.Columns}) {
				return
			}
		}
	}
}

func (t Matrix2D[T]) IterCells() iter.Seq2[Cursor2D, T] {
	return func(yield func(Cursor2D, T) bool) {
		for row := 0; row < t.Rows; row++ {
			for col := 0; col < t.Columns; col++ {
				cell := t.data[row][col]
				if !yield(Cursor2D{Row: row, Column: col}, cell) {
					return
				}
			}
		}
	}
}

func (t Matrix2D[T]) IterRowsReverse() iter.Seq2[int, Vector[T]] {
	return func(yield func(int, Vector[T]) bool) {
		for row := t.Rows - 1; row >= 0; row-- {
			arr := t.data[row]
			if !yield(row, Vector[T]{data: arr, size: t.Columns}) {
				return
			}
		}
	}
}

func (t Matrix2D[T]) IterColumns() iter.Seq2[int, Vector[T]] {
	return func(yield func(int, Vector[T]) bool) {
		for column := 0; column < t.Columns; column++ {
			arr := make([]T, t.Rows)

			for y := range t.Columns {
				arr[y] = t.data[y][column]
			}
			if !yield(column, Vector[T]{data: arr, size: t.Rows}) {
				return
			}
		}
	}
}

func (t Matrix2D[T]) IterColumnsReverse() iter.Seq2[int, Vector[T]] {
	return func(yield func(int, Vector[T]) bool) {
		for column := t.Columns - 1; column >= 0; column-- {
			arr := make([]T, t.Rows)

			for y := range t.Columns {
				arr = append(arr, t.data[y][column])
			}
			if !yield(column, Vector[T]{data: arr, size: t.Rows}) {
				return
			}
		}
	}
}

func (t Matrix2D[T]) Iter(axis int) iter.Seq2[int, Vector[T]] {
	if axis == 0 {
		return t.IterColumns()
	} else if axis == 1 {
		return t.IterRows()
	} else {
		panic("Invalid axis number for 2D matrix")
	}
}

func (t Matrix2D[T]) IterReverse(axis int) iter.Seq2[int, Vector[T]] {
	if axis == 0 {
		return t.IterColumnsReverse()
	} else if axis == 1 {
		return t.IterRowsReverse()
	} else {
		panic("Invalid axis number for 2D matrix")
	}
}

func (t Matrix2D[T]) SubMatrix(offsetRows, offsetCols, sizeRows, sizeCols int) Matrix2D[T] {
	if offsetRows < 0 || offsetCols < 0 || sizeRows < 0 || sizeCols < 0 || offsetRows+sizeRows > t.Rows || offsetCols+sizeCols > t.Columns {
		panic("out of bounds for SubMatrix")
	}

	matrix := make([][]T, sizeRows)
	for i := 0; i < sizeRows; i++ {
		matrix[i] = make([]T, sizeCols)
	}

	for row := 0; row < sizeRows; row++ {
		for col := 0; col < sizeCols; col++ {
			matrix[row][col] = t.data[row+offsetRows][col+offsetCols]
		}
	}
	return Matrix2D[T]{data: matrix, Rows: sizeRows, Columns: sizeCols}
}

func (t Matrix2D[T]) RotateRight() Matrix2D[T] {
	matrix := make([][]T, t.Columns)
	for i := 0; i < t.Rows; i++ {
		matrix[i] = make([]T, t.Rows)
	}

	for row := 0; row < t.Rows; row++ {
		for col := 0; col < t.Columns; col++ {
			matrix[col][-row-1+t.Rows] = t.data[row][col]
		}
	}
	return Matrix2D[T]{data: matrix, Rows: t.Columns, Columns: t.Rows}
}

func (t Matrix2D[T]) Get(row, col int) T {
	return t.data[row][col]
}

func (t Matrix2D[T]) Set(row, col int, value T) {
	t.data[row][col] = value
}

func (t Matrix2D[T]) CheckBoundary(row, col int) bool {
	if row < 0 || row >= t.Rows || col < 0 || col >= t.Columns {
		return false
	} else {
		return true
	}
}

func (t Matrix2D[T]) Clone() Matrix2D[T] {
	var data = make([][]T, t.Rows)

	for rowIdx, row := range t.IterRows() {
		newRow := make([]T, t.Columns)
		for colIdx, cell := range row.Iterator() {
			newRow[colIdx] = cell
		}
		data[rowIdx] = newRow
	}
	return Matrix2D[T]{data: data, Rows: t.Rows, Columns: t.Columns}
}

func (t Matrix2D[T]) Fill(item T) {
	for rowIdx, row := range t.IterRows() {
		for colIdx := range row.Iterator() {
			t.data[rowIdx][colIdx] = item
		}
	}
}

func (t Matrix2D[T]) SwapRows(rowA, rowB int) {
	temp := t.data[rowA]
	t.data[rowA] = t.data[rowB]
	t.data[rowB] = temp
}

func (t Matrix2D[T]) Print(mapper func(T) string) {
	for rowIdx, row := range t.IterRows() {
		for colIdx := range row.Iterator() {
			fmt.Printf(mapper(t.data[rowIdx][colIdx]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func (t Matrix2D[rune]) PrintChars() {
	t.Print(func(r rune) string {
		return fmt.Sprintf("%c", r)
	})
}
