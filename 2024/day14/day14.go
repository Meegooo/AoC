package main

import (
	"bufio"
	"fmt"
	"github.com/meegoue/AoC/library/maths"
	"github.com/meegoue/AoC/library/matrix"
	"io"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	res := part2(strings.NewReader("AAAAAA\nAAABBA\nAAABBA\nABBAAA\nABBAAA\nAAAAAA"))
	fmt.Println(res)
}

type robot struct {
	posX, posY, velX, velY int
}

func part1(reader io.Reader) string {
	parseRegex := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	fieldSizeText := strings.Split(scanner.Text(), " ")
	rows := maths.AtoiUnwrap(fieldSizeText[0])
	cols := maths.AtoiUnwrap(fieldSizeText[1])
	midRow := rows / 2
	midCol := cols / 2

	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0
	movedRobots := make([]robot, 0)
	for scanner.Scan() {
		line := scanner.Text()
		matches := parseRegex.FindStringSubmatch(line)
		r := robot{
			posX: maths.AtoiUnwrap(matches[1]),
			posY: maths.AtoiUnwrap(matches[2]),
			velX: maths.AtoiUnwrap(matches[3]),
			velY: maths.AtoiUnwrap(matches[4]),
		}
		movedRobot := robot{
			posX: maths.Mod(r.posX+r.velX*100, cols),
			posY: maths.Mod(r.posY+r.velY*100, rows),
			velX: r.velX,
			velY: r.velY,
		}
		movedRobots = append(movedRobots, movedRobot)
		if movedRobot.posX < midCol && movedRobot.posY < midRow {
			q1++
		}
		if movedRobot.posX > midCol && movedRobot.posY < midRow {
			q2++
		}
		if movedRobot.posX < midCol && movedRobot.posY > midRow {
			q3++
		}
		if movedRobot.posX > midCol && movedRobot.posY > midRow {
			q4++
		}
	}

	field := matrix.NewEmptyMatrix2D[rune](rows, cols)
	field.Fill('.')
	for _, r := range movedRobots {
		c := field.Get(r.posY, r.posX)
		if c == '.' {
			c = '1'
		} else {
			c++
		}
		field.Set(r.posY, r.posX, c)
	}
	field.Print(func(r rune) string {
		return fmt.Sprintf("%c", r)
	})
	result := q1 * q2 * q3 * q4
	return strconv.Itoa(result)
}

func part2(reader io.Reader) string {
	parseRegex := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	fieldSizeText := strings.Split(scanner.Text(), " ")
	rows := maths.AtoiUnwrap(fieldSizeText[0])
	cols := maths.AtoiUnwrap(fieldSizeText[1])
	if rows < 20 {
		return ""
	}

	robots := make([]robot, 0)
	for scanner.Scan() {
		line := scanner.Text()
		matches := parseRegex.FindStringSubmatch(line)
		r := robot{
			posX: maths.AtoiUnwrap(matches[1]),
			posY: maths.AtoiUnwrap(matches[2]),
			velX: maths.AtoiUnwrap(matches[3]),
			velY: maths.AtoiUnwrap(matches[4]),
		}
		robots = append(robots, r)
	}

	field := matrix.NewEmptyMatrix2D[rune](rows, cols)

	for iter := range 10000 {
		field.Fill('.')
		for i, r := range robots {
			c := field.Get(r.posY, r.posX)
			if c == '.' {
				c = '1'
			} else {
				c++
			}
			field.Set(r.posY, r.posX, c)
			movedRobot := robot{
				posX: maths.Mod(r.posX+r.velX, cols),
				posY: maths.Mod(r.posY+r.velY, rows),
				velX: r.velX,
				velY: r.velY,
			}
			robots[i] = movedRobot
		}

		bad := false
		for _, i := range field.IterCells() {
			if i != '.' && i != '1' {
				bad = true
			}
		}
		if !bad || iter%103 == 0 {
			fmt.Println(iter)
			field.Print(func(r rune) string {
				return fmt.Sprintf("%c", r)
			})
			if !bad {
				break
			}
		}

	}
	return ""
}
