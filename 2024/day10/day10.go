package main

import (
	"github.com/meegoue/AoC/library/matrix"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day9/test.txt")
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	part1(strings.NewReader("09091"))
}

func part1(reader io.Reader) string {
	mtx := matrix.ReadMatrix2D(reader)

	counter := 0
	for cursor, cell := range mtx.IterCells() {
		if cell == '0' {
			visited := mtx.Clone()
			visited.Fill(' ')
			runIterPart1(&mtx, &visited, cursor)
			for _, visitedCell := range visited.IterCells() {
				if visitedCell == '9' {
					counter++
				}
			}
			//visited.Print(func(r rune) string {
			//	return fmt.Sprintf("%c", r)
			//})
		}
	}
	return strconv.Itoa(counter)
}

func part2(reader io.Reader) string {
	mtx := matrix.ReadMatrix2D(reader)

	counter := 0
	for cursor, cell := range mtx.IterCells() {
		if cell == '0' {
			visited := mtx.Clone()
			visited.Fill(' ')
			counter += runIterPart2(&mtx, &visited, cursor)
			//visited.Print(func(r rune) string {
			//	return fmt.Sprintf("%c", r)
			//})
		}
	}
	return strconv.Itoa(counter)
}

func runIterPart1(mtx *matrix.Matrix2D[rune], visited *matrix.Matrix2D[rune], start matrix.Cursor2D) {
	currentValue := mtx.Get(start.Row, start.Column)
	visited.Set(start.Row, start.Column, currentValue)
	wantedValue := currentValue + 1

	current := start
	for range 4 {
		next := current.Forward()
		if mtx.CheckBoundary(next.Row, next.Column) && visited.Get(next.Row, next.Column) == ' ' && mtx.Get(next.Row, next.Column) == wantedValue {
			runIterPart1(mtx, visited, next)
		}
		current = current.TurnClockwise()
	}
}

func runIterPart2(mtx *matrix.Matrix2D[rune], visited *matrix.Matrix2D[rune], start matrix.Cursor2D) int {
	currentValue := mtx.Get(start.Row, start.Column)
	if currentValue == '9' {
		return 1
	}
	visited.Set(start.Row, start.Column, currentValue)
	wantedValue := currentValue + 1
	counter := 0
	current := start
	for range 4 {
		next := current.Forward()
		if mtx.CheckBoundary(next.Row, next.Column) && mtx.Get(next.Row, next.Column) == wantedValue {
			counter += runIterPart2(mtx, visited, next)
		}
		current = current.TurnClockwise()
	}
	return counter
}
