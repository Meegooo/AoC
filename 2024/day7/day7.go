package main

import (
	"bufio"
	"fmt"
	"github.com/meegoue/AoC/library/collections"
	"github.com/meegoue/AoC/library/maths"
	"io"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
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
	scanner := bufio.NewScanner(reader)
	re := regexp.MustCompile(`(\d+): ((?: ?\d+)+)`)

	accumulator := 0

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		result, _ := strconv.Atoi(matches[1])
		elements := slices.Collect(collections.Map(slices.Values(strings.Split(matches[2], " ")), maths.AtoiUnwrap))
		if rec1(elements, 1, int64(elements[0]), int64(result)) {
			fmt.Println(line)
			accumulator += result
		}
	}
	return strconv.Itoa(accumulator)
}

func part2(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	re := regexp.MustCompile(`(\d+): ((?: ?\d+)+)`)

	accumulator := 0

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		result, _ := strconv.Atoi(matches[1])
		elements := slices.Collect(collections.Map(slices.Values(strings.Split(matches[2], " ")), maths.AtoiUnwrap))
		if rec2(elements, 1, int64(elements[0]), int64(result)) {
			fmt.Println(line)
			accumulator += result
		}
	}
	return strconv.Itoa(accumulator)
}

func rec1(elements []int, idx int, accumulator int64, target int64) bool {
	if accumulator > target {
		return false
	}
	if idx == len(elements) {
		return accumulator == target
	}
	return rec1(elements, idx+1, accumulator+int64(elements[idx]), target) ||
		rec1(elements, idx+1, accumulator*int64(elements[idx]), target)
}

func rec2(elements []int, idx int, accumulator int64, target int64) bool {
	if accumulator > target {
		return false
	}
	if idx == len(elements) {
		return accumulator == target
	}
	concat, _ := strconv.ParseInt(strconv.FormatInt(accumulator, 10)+strconv.FormatInt(int64(elements[idx]), 10), 10, 64)
	return rec2(elements, idx+1, accumulator+int64(elements[idx]), target) ||
		rec2(elements, idx+1, accumulator*int64(elements[idx]), target) ||
		rec2(elements, idx+1, concat, target)
}
