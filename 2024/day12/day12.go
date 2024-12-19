package main

import (
	"fmt"
	"github.com/gammazero/deque"
	"github.com/meegoue/AoC/library/matrix"
	"io"
	"strconv"
	"strings"
)

func main() {
	res := part2(strings.NewReader("AAAAAA\nAAABBA\nAAABBA\nABBAAA\nABBAAA\nAAAAAA"))
	fmt.Println(res)
}

func part1(reader io.Reader) string {
	mtx := matrix.ReadMatrix2D(reader)
	visited := mtx.Clone()
	visited.Fill(' ')
	totalCost := 0
	zoneId := ' ' + 1
	for coords, _ := range mtx.IterCells() {
		if visited.Get(coords.Row, coords.Column) != ' ' {
			continue
		}
		perimeter, area := calculateAndFill(&mtx, &visited, coords, zoneId)
		totalCost += area * perimeter
		zoneId += 1
	}
	return strconv.Itoa(totalCost)
}

func part2(reader io.Reader) string {
	mtx := matrix.ReadMatrix2D(reader)
	visited := mtx.Clone()
	visited.Fill(' ')
	zoneId := ' ' + 1
	var areas = make(map[rune]int)
	var sides = make(map[rune]int)
	for coords, _ := range mtx.IterCells() {
		if visited.Get(coords.Row, coords.Column) != ' ' {
			continue
		}
		_, area := calculateAndFill(&mtx, &visited, coords, zoneId)
		areas[zoneId] = area
		zoneId += 1
	}

	for range 4 {
		prev := ' '
		for colIdx, _ := range visited.IterColumns() {
			cell := visited.Get(0, colIdx)
			if prev != cell {
				sides[cell] += 1
				prev = cell
			}
		}

		for rowIdx, row := range visited.IterRows() {
			if rowIdx == mtx.Rows-1 {
				continue
			}
			prev = ' '
			for colIdx, upCell := range row.Iterator() {
				downCell := visited.Get(rowIdx+1, colIdx)
				if upCell == downCell {
					prev = ' '
				} else if upCell != prev {
					sides[upCell] += 1
					prev = upCell
				}
			}
		}
		visited = visited.RotateRight()
	}

	totalCost := 0
	for key, value := range sides {
		totalCost += areas[key] * value
	}
	visited.Print(func(r rune) string {
		return fmt.Sprintf("%c", r)
	})
	return strconv.Itoa(totalCost)
}

func calculateAndFill(mtx *matrix.Matrix2D[rune], visited *matrix.Matrix2D[rune], startFrom matrix.Cursor2D, zoneId rune) (int, int) {
	var queue deque.Deque[matrix.Cursor2D]
	queue.PushBack(startFrom)
	targetRune := mtx.Get(startFrom.Row, startFrom.Column)
	perimeter := 0
	area := 0
	for queue.Len() != 0 {
		elem := queue.PopFront()
		if visited.Get(elem.Row, elem.Column) != ' ' {
			continue
		}
		rotatedElem := elem.TurnUp()
		visited.Set(elem.Row, elem.Column, zoneId)
		area += 1
		currentPerimeter := 4
		for range 4 {
			// Check around
			elemCopy := rotatedElem.Forward()
			if mtx.CheckBoundary(elemCopy.Row, elemCopy.Column) && mtx.Get(elemCopy.Row, elemCopy.Column) == targetRune {
				if visited.Get(elemCopy.Row, elemCopy.Column) == ' ' {
					queue.PushBack(elemCopy)
				}
				currentPerimeter -= 1
			}
			rotatedElem = rotatedElem.TurnClockwise()
		}
		perimeter += currentPerimeter
	}
	return perimeter, area
}
