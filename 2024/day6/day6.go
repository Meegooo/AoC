package main

import (
	"bufio"
	"github.com/meegoue/AoC/library/matrix"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day6/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	part1(file)
}

func part1(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	var lines [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, []rune(line))
	}
	var mtx = matrix.NewMatrix2D(lines)
	var cursor matrix.Cursor2D
	for rowIdx, row := range mtx.IterRows() {
		for colIdx, cell := range row.Iterator(0) {
			if cell == '^' {
				cursor = matrix.Cursor2D{Row: rowIdx, Column: colIdx, Direction: matrix.UP}
			}
		}
	}

	var mtxClone = mtx.Clone()
	mtxClone.Set(cursor.Row, cursor.Column, 'X')

	for {
		next := cursor.Forward()
		if mtx.CheckBoundary(next.Row, next.Column) {
			if mtx.Get(next.Row, next.Column) == '#' {
				cursor = cursor.TurnClockwise()
			} else {
				cursor = next
				mtxClone.Set(cursor.Row, cursor.Column, 'X')
				//mtxClone.Print(func(r rune) string {
				//	return fmt.Sprintf("%c", r)
				//})
			}
		} else {
			count := 0
			for _, cell := range mtxClone.IterCells() {
				if cell == 'X' {
					count += 1
				}
			}
			return strconv.Itoa(count)
		}
	}
}

func part2(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	var lines [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, []rune(line))
	}
	var mtx = matrix.NewMatrix2D(lines)
	var cursorStart matrix.Cursor2D
	for rowIdx, row := range mtx.IterRows() {
		for colIdx, cell := range row.Iterator(0) {
			if cell == '^' {
				cursorStart = matrix.Cursor2D{Row: rowIdx, Column: colIdx, Direction: matrix.UP}
			}
		}
	}

	ch := make(chan bool, 100000)
	for coords := range mtx.IterCells() {
		go runIter(&mtx, &coords, &cursorStart, ch)
	}

	counter := 0
	for range mtx.IterCells() {
		res := <-ch
		if res {
			counter++
		}
	}
	return strconv.Itoa(counter)
}

func runIter(mtx *matrix.Matrix2D[rune], coords *matrix.Cursor2D, startPoint *matrix.Cursor2D, ch chan bool) {
	cellValue := mtx.Get(coords.Row, coords.Column)
	if cellValue != '.' {
		ch <- false
		return
	}
	currentMtx := mtx.Clone()
	currentMtx.Set(coords.Row, coords.Column, '#')

	cursors := make(map[matrix.Cursor2D]bool)
	cursor := startPoint.Clone()
	for {
		next := cursor.Forward()
		if currentMtx.CheckBoundary(next.Row, next.Column) {
			if currentMtx.Get(next.Row, next.Column) == '#' {
				cursor = cursor.TurnClockwise()
			} else if cursors[next] == true {
				ch <- true
				return
			} else {
				cursor = next
			}
			cursors[cursor] = true
		} else {
			ch <- false
			return
		}
	}
}
