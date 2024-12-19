package main

import (
	"fmt"
	"github.com/meegoue/AoC/library/matrix"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day7/test.txt")
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	part1(file)
}

func part1(reader io.Reader) string {
	mtx := matrix.ReadMatrix2D(reader)
	coordsByFrequencies := make(map[rune][]matrix.Cursor2D)
	for cursor, cell := range mtx.IterCells() {
		if cell != '.' {
			coordsByFrequencies[cell] = append(coordsByFrequencies[cell], cursor)
		}
	}

	nodes := make(map[matrix.Cursor2D]bool)

	for key := range coordsByFrequencies {
		for idx1, first := range coordsByFrequencies[key] {
			for idx2, second := range coordsByFrequencies[key] {
				if idx2 <= idx1 {
					continue
				}
				diff := first.ToVector().Subtract(second.ToVector())
				result1 := first.ToVector().Add(diff)
				result2 := second.ToVector().Add(diff.Invert())

				nodes[result1.ToCursor2D()] = true
				nodes[result2.ToCursor2D()] = true
			}
		}
	}
	counter := 0
	for node := range nodes {
		if node.Row >= 0 && node.Row < mtx.Rows && node.Column >= 0 && node.Column < mtx.Columns {
			counter++
		}
	}

	return strconv.Itoa(counter)
}

func part2(reader io.Reader) string {
	mtx := matrix.ReadMatrix2D(reader)
	coordsByFrequencies := make(map[rune][]matrix.Cursor2D)
	for cursor, cell := range mtx.IterCells() {
		if cell != '.' {
			coordsByFrequencies[cell] = append(coordsByFrequencies[cell], cursor)
		}
	}

	nodes := make(map[matrix.Cursor2D]bool)

	counter := 0

	for key := range coordsByFrequencies {
		for idx1, first := range coordsByFrequencies[key] {
			for idx2, second := range coordsByFrequencies[key] {
				if idx2 <= idx1 {
					continue
				}
				diff := first.ToVector().Subtract(second.ToVector())
				current := first.ToVector()
				for {
					currentCursor := current.ToCursor2D()
					if currentCursor.Row >= 0 && currentCursor.Row < mtx.Rows && currentCursor.Column >= 0 && currentCursor.Column < mtx.Columns {
						nodes[currentCursor] = true
						current = current.Add(diff)
					} else {
						break
					}
				}
				current = second.ToVector()
				for {
					currentCursor := current.ToCursor2D()
					if currentCursor.Row >= 0 && currentCursor.Row < mtx.Rows && currentCursor.Column >= 0 && currentCursor.Column < mtx.Columns {
						nodes[currentCursor] = true
						current = current.Subtract(diff)
					} else {
						break
					}
				}
			}
		}
	}

	printMtx := mtx.Clone()
	for node := range nodes {
		counter++
		printMtx.Set(node.Row, node.Column, '#')
	}
	printMtx.Print(func(r rune) string {
		return fmt.Sprintf("%c", r)
	})

	return strconv.Itoa(counter)
}
