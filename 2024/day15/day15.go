package main

import (
	"bufio"
	"fmt"
	"github.com/meegoue/AoC/library/matrix"
	"io"
	"strconv"
	"strings"
)

func main() {
	res := part1(strings.NewReader("########\n#..O.O.#\n##@.O..#\n#...O..#\n#.#.O..#\n#...O..#\n#......#\n########\n\n<^^>>>vv<v>>v<<"))
	fmt.Println(res)
}

func part1(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	var lines [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, []rune(line))
	}
	field := matrix.NewMatrix2D(lines)

	var moves []rune
	for scanner.Scan() {
		line := []rune(scanner.Text())
		moves = append(moves, line...)
	}

	var robotPosition matrix.Cursor2D
	for cur, r := range field.IterCells() {
		if r == '@' {
			robotPosition = cur.TurnUp()
			field.Set(cur.Row, cur.Column, '.')
			break
		}
	}

	for _, move := range moves {
		//field.PrintChars()
		pointedRobotCursor := robotPosition
		switch move {
		case '>':
			pointedRobotCursor = pointedRobotCursor.TurnRight()
		case '^':
			pointedRobotCursor = pointedRobotCursor.TurnUp()
		case '<':
			pointedRobotCursor = pointedRobotCursor.TurnLeft()
		case 'v':
			pointedRobotCursor = pointedRobotCursor.TurnDown()
		}

		chainEnd := pointedRobotCursor
		hitWall := false
		for {
			chainEnd = chainEnd.Forward()
			inspectedCell := field.Get(chainEnd.Row, chainEnd.Column)
			if inspectedCell == '#' {
				hitWall = true
				break
			} else if inspectedCell == '.' {
				break
			}
		}
		if hitWall {
			continue
		}
		// Turn around
		chainEnd = chainEnd.TurnClockwise()
		chainEnd = chainEnd.TurnClockwise()

		// Pull stuff, until we reach pointedRobotCursor
		for {
			if chainEnd.Row == pointedRobotCursor.Row && chainEnd.Column == pointedRobotCursor.Column {
				break
			}
			pullFrom := chainEnd.Forward()
			field.Set(chainEnd.Row, chainEnd.Column, field.Get(pullFrom.Row, pullFrom.Column))
			field.Set(pullFrom.Row, pullFrom.Column, '.')
			chainEnd = pullFrom
		}

		robotPosition = pointedRobotCursor.Forward()
	}
	field.PrintChars()

	gps := 0
	for cursor, cell := range field.IterCells() {
		if cell == 'O' {
			gps += cursor.Row*100 + cursor.Column
		}
	}

	return strconv.Itoa(gps)
}

func part2(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	var lines [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		var newLine = make([]rune, 0)
		for _, r := range line {
			switch r {
			case '#':
				newLine = append(newLine, '#', '#')
			case 'O':
				newLine = append(newLine, '[', ']')
			case '.':
				newLine = append(newLine, '.', '.')
			case '@':
				newLine = append(newLine, '@', '.')
			}
		}
		lines = append(lines, newLine)
	}
	field := matrix.NewMatrix2D(lines)

	var moves []rune
	for scanner.Scan() {
		line := []rune(scanner.Text())
		moves = append(moves, line...)
	}

	var robotPosition matrix.Cursor2D
	for cur, r := range field.IterCells() {
		if r == '@' {
			robotPosition = cur.TurnUp()
			field.Set(cur.Row, cur.Column, '.')
			break
		}
	}

	for _, move := range moves {
		//field.PrintChars()
		pointedRobotCursor := robotPosition
		switch move {
		case '>':
			pointedRobotCursor = pointedRobotCursor.TurnRight()
		case '^':
			pointedRobotCursor = pointedRobotCursor.TurnUp()
		case '<':
			pointedRobotCursor = pointedRobotCursor.TurnLeft()
		case 'v':
			pointedRobotCursor = pointedRobotCursor.TurnDown()
		}
		robotPosition = push(pointedRobotCursor, &field)
	}
	field.PrintChars()

	gps := 0
	for cursor, cell := range field.IterCells() {
		if cell == '[' {
			gps1 := cursor.Row*100 + cursor.Column
			rightCursor := cursor.MoveRight()
			gps2 := rightCursor.Row*100 + rightCursor.Column
			gps += min(gps1, gps2)
		}
	}

	return strconv.Itoa(gps)
}

func tryPush(source matrix.Cursor2D, field *matrix.Matrix2D[rune]) bool {
	moved := source.Forward()
	cellUnderMoved := field.Get(moved.Row, moved.Column)
	if cellUnderMoved == '#' {
		return false
	} else if cellUnderMoved == '.' {
		return true
	} else if cellUnderMoved == '[' {
		if source.Direction == matrix.RIGHT {
			movedTwice := moved.Forward()
			return tryPush(movedTwice, field)
		} else if source.Direction == matrix.UP || source.Direction == matrix.DOWN {
			tryPushLeft := tryPush(moved, field)
			tryPushRight := tryPush(moved.MoveRight(), field)
			return tryPushLeft && tryPushRight
		}
	} else if cellUnderMoved == ']' {
		if source.Direction == matrix.LEFT {
			movedTwice := moved.Forward()
			return tryPush(movedTwice, field)
		} else if source.Direction == matrix.UP || source.Direction == matrix.DOWN {
			tryPushLeft := tryPush(moved.MoveLeft(), field)
			tryPushRight := tryPush(moved, field)
			return tryPushLeft && tryPushRight
		}
	}
	panic("Unexpected cell")
}

func push(source matrix.Cursor2D, field *matrix.Matrix2D[rune]) matrix.Cursor2D {
	pushPossible := tryPush(source, field)
	if !pushPossible {
		return source
	}
	moved := source.Forward()
	movedTwice := moved.Forward()
	cellUnderMoved := field.Get(moved.Row, moved.Column)
	if cellUnderMoved == '#' {
		panic("Impossible")
	} else if cellUnderMoved == '.' {
		return source.Forward()
	} else if cellUnderMoved == '[' {
		if source.Direction == matrix.RIGHT {
			movedThrice := movedTwice.Forward()
			push(movedTwice, field)
			field.Set(moved.Row, moved.Column, '.')
			field.Set(movedTwice.Row, movedTwice.Column, '[')
			field.Set(movedThrice.Row, movedThrice.Column, ']')
			return moved
		} else if source.Direction == matrix.UP || source.Direction == matrix.DOWN {
			movedRight := moved.MoveRight()
			push(moved, field)
			push(movedRight, field)
			field.Set(moved.Row, moved.Column, '.')
			field.Set(movedRight.Row, movedRight.Column, '.')
			field.Set(movedTwice.Row, movedTwice.Column, '[')
			field.Set(movedRight.Forward().Row, movedRight.Forward().Column, ']')
		}
		return moved
	} else if cellUnderMoved == ']' {
		if source.Direction == matrix.LEFT {
			movedTwice := moved.Forward()
			movedThrice := movedTwice.Forward()
			push(movedTwice, field)
			field.Set(moved.Row, moved.Column, '.')
			field.Set(movedTwice.Row, movedTwice.Column, ']')
			field.Set(movedThrice.Row, movedThrice.Column, '[')
		} else if source.Direction == matrix.UP || source.Direction == matrix.DOWN {
			movedLeft := moved.MoveLeft()
			push(moved, field)
			push(movedLeft, field)
			field.Set(moved.Row, moved.Column, '.')
			field.Set(movedLeft.Row, movedLeft.Column, '.')
			field.Set(movedTwice.Row, movedTwice.Column, ']')
			field.Set(movedLeft.Forward().Row, movedLeft.Forward().Column, '[')
		}
		return moved
	}
	panic("Unexpected cell")
}
