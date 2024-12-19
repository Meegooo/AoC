package matrix

import (
	"github.com/meegoue/AoC/library/collections"
	"golang.org/x/exp/constraints"
	"iter"
	"slices"
)

type Vector[T constraints.Integer | constraints.Float] struct {
	data []T
	size int
}

func (t Vector[T]) Iterator() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for row := 0; row < t.size; row++ {
			if !yield(row, t.data[row]) {
				return
			}
		}
	}
}

func (t Vector[T]) IteratorReverse() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for row := t.size - 1; row >= 0; row-- {
			if !yield(row, t.data[row]) {
				return
			}
		}
	}
}

func (t Vector[T]) Invert() Vector[T] {
	inverted := slices.Collect(collections.Map(slices.Values(t.data), func(t T) T { return -t }))
	return Vector[T]{data: inverted, size: t.size}
}

func (t Vector[T]) Get(idx int) T {
	return t.data[idx]
}

func (t Vector[T]) Subtract(other Vector[T]) Vector[T] {
	if t.size != other.size {
		panic("Trying to subtract vectors of different sizes")
	}

	result := make([]T, t.size)
	for i := range t.Iterator() {
		l := t.data[i]
		r := other.data[i]
		result[i] = l - r
	}
	return Vector[T]{result, t.size}
}

func (t Vector[T]) Add(other Vector[T]) Vector[T] {
	if t.size != other.size {
		panic("Trying to subtract vectors of different sizes")
	}

	result := make([]T, t.size)
	for i := range t.Iterator() {
		l := t.data[i]
		r := other.data[i]
		result[i] = l + r
	}
	return Vector[T]{result, t.size}
}

func (t Vector[T]) ToCursor2D() Cursor2D {
	if t.size != 2 {
		panic("Trying to convert vector with size other than 2 to Cursor2D")
	}
	row, okRow := any(t.data[0]).(int)
	col, okCol := any(t.data[1]).(int)
	if okRow && okCol {
		return Cursor2D{Row: row, Column: col}
	}
	panic("Vector of non-int type can't be converted to cursor")
}
