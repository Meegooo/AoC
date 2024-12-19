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
)

func main() {
	file, err := os.Open("day3/actual.txt")
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
	var result int64 = 0
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		for matchIdx := range matches {
			match := matches[matchIdx]
			ints := slices.Collect(collections.Map(collections.Map(collections.SkipOne(slices.Values(match), 0), maths.AtoiUnwrap), func(t int) int64 {
				return int64(t)
			}))
			if len(ints) != 2 {
				fmt.Println(ints)
			}
			result += ints[0] * ints[1]
		}
	}
	return strconv.FormatInt(result, 10)
}

func part2(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	var result int64 = 0
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	enabled := true
	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		for matchIdx := range matches {
			match := matches[matchIdx]
			if match[0] == "don't()" {
				enabled = false
			} else if match[0] == "do()" {
				enabled = true
			} else if enabled {
				ints := slices.Collect(collections.Map(collections.Map(collections.SkipOne(slices.Values(match), 0), maths.AtoiUnwrap), func(t int) int64 {
					return int64(t)
				}))
				if len(ints) != 2 {
					fmt.Println(ints)
				}
				result += ints[0] * ints[1]
			}

		}
	}
	return strconv.FormatInt(result, 10)
}
