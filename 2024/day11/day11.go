package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	arr := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	a := arr[:4]
	b := arr[4:]
	println(a, b)
	file, err := os.Open("day11/test.txt")
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	part1(strings.NewReader("09091"))
}

func part1(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	res := 0
	for _, stone := range input {
		res += runIterPart2(stone, 25)
	}
	return strconv.Itoa(res)
}
func part2(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	res := 0
	for _, stone := range input {
		res += runIterPart2(stone, 75)
	}
	return strconv.Itoa(res)
}
func runIterPart1(input []string) []string {
	out := make([]string, 0)
	for _, stone := range input {
		if stone == "0" {
			out = append(out, "1")
		} else if len(stone)%2 == 0 {
			mid := len(stone) / 2
			out = append(out, removeLeadingZeros(stone[:mid]), removeLeadingZeros(stone[mid:]))
		} else {
			i, _ := strconv.ParseInt(stone, 10, 64)
			out = append(out, strconv.FormatInt(i*2024, 10))
		}
	}
	return out
}

type k struct {
	value string
	iter  int
}

var cache = make(map[k]int)

func runIterPart2(stone string, iter int) int {
	if iter == 0 {
		return 1
	}
	key := k{value: stone, iter: iter}
	value, ok := cache[key]
	if ok {
		return value
	}

	if stone == "0" {
		res := runIterPart2("1", iter-1)
		cache[key] = res
		return res
	} else if len(stone)%2 == 0 {
		mid := len(stone) / 2
		left := removeLeadingZeros(stone[:mid])
		right := removeLeadingZeros(stone[mid:])
		res := runIterPart2(left, iter-1) + runIterPart2(right, iter-1)
		cache[key] = res
		return res
	} else {
		i, _ := strconv.ParseInt(stone, 10, 64)
		res := runIterPart2(strconv.FormatInt(i*2024, 10), iter-1)
		cache[key] = res
		return res
	}
}

var leadingZeros = regexp.MustCompile(`^0*`)

func removeLeadingZeros(input string) string {
	r := leadingZeros.ReplaceAllLiteralString(input, "")
	if r == "" {
		return "0"
	} else {
		return r
	}
}
