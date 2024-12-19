package main

import (
	"bufio"
	"fmt"
	"github.com/meegoue/AoC/library/matrix"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day4/actual.txt")
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
	var cloneMtx = make([][]rune, mtx.Rows)
	for idx, _ := range mtx.IterRows() {
		cloneMtx[idx] = make([]rune, mtx.Columns)
	}

	instances := 0
	for rowIdx, row := range mtx.IterRows() {
		for colIdx := 0; colIdx < mtx.Columns-3; colIdx++ {
			if rowIdx+3 < mtx.Rows {
				word := fmt.Sprintf("%c%c%c%c", mtx.Get(rowIdx, colIdx), mtx.Get(rowIdx+1, colIdx+1), mtx.Get(rowIdx+2, colIdx+2), mtx.Get(rowIdx+3, colIdx+3))
				if word == "XMAS" || word == "SAMX" {
					cloneMtx[rowIdx][colIdx] = mtx.Get(rowIdx, colIdx)
					cloneMtx[rowIdx+1][colIdx+1] = mtx.Get(rowIdx+1, colIdx+1)
					cloneMtx[rowIdx+2][colIdx+2] = mtx.Get(rowIdx+2, colIdx+2)
					cloneMtx[rowIdx+3][colIdx+3] = mtx.Get(rowIdx+3, colIdx+3)
					instances++
				}
			}
			if rowIdx >= 3 {
				word := fmt.Sprintf("%c%c%c%c", mtx.Get(rowIdx, colIdx), mtx.Get(rowIdx-1, colIdx+1), mtx.Get(rowIdx-2, colIdx+2), mtx.Get(rowIdx-3, colIdx+3))
				if word == "XMAS" || word == "SAMX" {
					cloneMtx[rowIdx][colIdx] = mtx.Get(rowIdx, colIdx)
					cloneMtx[rowIdx-1][colIdx+1] = mtx.Get(rowIdx-1, colIdx+1)
					cloneMtx[rowIdx-2][colIdx+2] = mtx.Get(rowIdx-2, colIdx+2)
					cloneMtx[rowIdx-3][colIdx+3] = mtx.Get(rowIdx-3, colIdx+3)
					instances++
				}
			}
			word := fmt.Sprintf("%c%c%c%c", row.Get(colIdx), row.Get(colIdx+1), row.Get(colIdx+2), row.Get(colIdx+3))
			if word == "XMAS" || word == "SAMX" {
				cloneMtx[rowIdx][colIdx] = mtx.Get(rowIdx, colIdx)
				cloneMtx[rowIdx][colIdx+1] = mtx.Get(rowIdx, colIdx+1)
				cloneMtx[rowIdx][colIdx+2] = mtx.Get(rowIdx, colIdx+2)
				cloneMtx[rowIdx][colIdx+3] = mtx.Get(rowIdx, colIdx+3)
				instances++
			}
		}
	}
	for colIdx, column := range mtx.IterColumns() {
		for rowIdx := 0; rowIdx < mtx.Rows-3; rowIdx++ {
			word := fmt.Sprintf("%c%c%c%c", column.Get(rowIdx), column.Get(rowIdx+1), column.Get(rowIdx+2), column.Get(rowIdx+3))
			if word == "XMAS" || word == "SAMX" {
				cloneMtx[rowIdx][colIdx] = mtx.Get(rowIdx, colIdx)
				cloneMtx[rowIdx+1][colIdx] = mtx.Get(rowIdx+1, colIdx)
				cloneMtx[rowIdx+2][colIdx] = mtx.Get(rowIdx+2, colIdx)
				cloneMtx[rowIdx+3][colIdx] = mtx.Get(rowIdx+3, colIdx)
				instances++
			}
		}
	}
	//
	//for i := range cloneMtx {
	//	for j := range cloneMtx[i] {
	//		c := cloneMtx[i][j]
	//		if c != 0 {
	//			fmt.Printf("%c", c)
	//		} else {
	//			fmt.Print(".")
	//		}
	//	}
	//	fmt.Println()
	//}
	return strconv.Itoa(instances)
}

func part2(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	var lines [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, []rune(line))
	}
	instances := 0

	var mtx = matrix.NewMatrix2D(lines)
	var cloneMtx = make([][]rune, mtx.Rows)
	for idx, _ := range mtx.IterRows() {
		cloneMtx[idx] = make([]rune, mtx.Columns)
	}

	for row := 0; row < mtx.Rows-2; row++ {
		for col := 0; col < mtx.Columns-2; col++ {
			subMtx := mtx.SubMatrix(row, col, 3, 3)
			for range 4 {
				if subMtx.Get(0, 0) == 'M' &&
					subMtx.Get(0, 2) == 'S' &&
					subMtx.Get(1, 1) == 'A' &&
					subMtx.Get(2, 0) == 'M' &&
					subMtx.Get(2, 2) == 'S' {
					cloneMtx[row][col] = subMtx.Get(0, 0)
					cloneMtx[row+1][col+1] = subMtx.Get(1, 1)
					cloneMtx[row+2][col] = subMtx.Get(2, 0)
					cloneMtx[row][col+2] = subMtx.Get(0, 2)
					cloneMtx[row+2][col+2] = subMtx.Get(2, 2)
					instances += 1
				}
				subMtx = subMtx.RotateRight()
			}
		}
	}
	for i := range cloneMtx {
		for j := range cloneMtx[i] {
			c := cloneMtx[i][j]
			if c != 0 {
				fmt.Printf("%c", c)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	return strconv.Itoa(instances)
}
