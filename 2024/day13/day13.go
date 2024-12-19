package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/meegoue/AoC/library/maths"
	"github.com/meegoue/AoC/library/matrix"
	"io"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//go:embed test.txt
var test string

func main() {
	arr := [][]float64{{3.0, 2.0, -4.0, 3.0},
		{2.0, 3.0, 3.0, 15.0},
		{5.0, -3, 1.0, 14.0},
	}
	mtx := matrix.NewMatrix2D[float64](arr)
	matrix.ForwardElim(mtx)
	solution := matrix.BackSub(mtx)
	println(solution)
	buttonRegex := regexp.MustCompile(`(abc)(def)`)
	matches := buttonRegex.FindStringSubmatch("abcdef")
	for i := range matches {
		fmt.Println(matches[i])
	}

	res := part2(strings.NewReader(test))
	fmt.Println(res)
}

func part1_brute(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	buttonRegex := regexp.MustCompile(`Button .: X\+(\d+), Y\+(\d+)`)
	prizeRegex := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)
	cost := 0
	for scanner.Scan() {
		aLine := buttonRegex.FindStringSubmatch(scanner.Text())
		scanner.Scan()
		bLine := buttonRegex.FindStringSubmatch(scanner.Text())
		scanner.Scan()
		prizeLine := prizeRegex.FindStringSubmatch(scanner.Text())
		scanner.Scan()
		aX, _ := strconv.Atoi(aLine[1])
		aY, _ := strconv.Atoi(aLine[2])
		bX, _ := strconv.Atoi(bLine[1])
		bY, _ := strconv.Atoi(bLine[2])
		prizeX, _ := strconv.Atoi(prizeLine[1])
		prizeY, _ := strconv.Atoi(prizeLine[2])

		minCost := math.MaxInt64
		for a := range 100 {
			for b := range 100 {
				if aX*a+bX*b > prizeX || aY*a+bY*b > prizeY {
					break
				} else if aX*a+bX*b == prizeX && aY*a+bY*b == prizeY {
					if minCost != math.MaxInt64 {
						panic("aaa")
					}
					minCost = min(a*3+b*1, minCost)
					break
				}
			}
		}
		if minCost != math.MaxInt64 {
			println(minCost)
			cost += minCost
		}

	}
	return strconv.Itoa(cost)
}

func part1(reader io.Reader) string {
	res := solve(reader, 0, 100)
	return strconv.FormatInt(res, 10)
}

func part2(reader io.Reader) string {
	res := solve(reader, 10000000000000, math.MaxInt64)
	return strconv.FormatInt(res, 10)
}

func solve(reader io.Reader, offset float64, max int64) int64 {
	scanner := bufio.NewScanner(reader)
	buttonRegex := regexp.MustCompile(`Button .: X\+(\d+), Y\+(\d+)`)
	prizeRegex := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)
	cost := int64(0)
	for scanner.Scan() {
		aLine := buttonRegex.FindStringSubmatch(scanner.Text())
		scanner.Scan()
		bLine := buttonRegex.FindStringSubmatch(scanner.Text())
		scanner.Scan()
		prizeLine := prizeRegex.FindStringSubmatch(scanner.Text())
		scanner.Scan()
		augMtx := matrix.NewEmptyMatrix2D[float64](2, 3)
		augMtx.Set(0, 0, float64(maths.AtoiUnwrap(aLine[1])))
		augMtx.Set(0, 1, float64(maths.AtoiUnwrap(bLine[1])))
		augMtx.Set(0, 2, float64(maths.AtoiUnwrap(prizeLine[1])))
		augMtx.Set(1, 0, float64(maths.AtoiUnwrap(aLine[2])))
		augMtx.Set(1, 1, float64(maths.AtoiUnwrap(bLine[2])))
		augMtx.Set(1, 2, float64(maths.AtoiUnwrap(prizeLine[2])))
		originalMtx := augMtx.Clone()
		originalMtx.IterRowsReverse()
		augMtx.Set(0, 2, augMtx.Get(0, 2)+offset)
		augMtx.Set(1, 2, augMtx.Get(1, 2)+offset)

		singular := matrix.ForwardElim(augMtx)
		if singular != -1 {
			panic("AAAAA")
		}
		solution := matrix.BackSub(augMtx)
		eps := 1e-3
		if math.Abs(solution[0]-math.Round(solution[0])) < eps && math.Abs(solution[1]-math.Round(solution[1])) < eps {
			roundSolution := []int64{int64(math.Round(solution[0])), int64(math.Round(solution[1]))}
			if roundSolution[0] > 0 && roundSolution[0] <= max && roundSolution[1] > 0 && roundSolution[1] <= max {
				c := roundSolution[0]*3 + roundSolution[1]
				fmt.Println(c)
				cost += c
			}
		}
		// A*aX + B*bX = prizeX
		// A*aY + B*bY = prizeY
		// A=(prizeX - B*bX)/aX
		// (prizeX - B*bX)/aX * aY + B*bY = prizeY
		// aY*prizeX/aX - aY * B * bX/aX +
	}
	return cost
}
