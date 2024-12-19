package main

import (
	"bufio"
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
	file, err := os.Open("day5/actual.txt")
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
	rules := make(map[int][]int)
	scanner := bufio.NewScanner(reader)
	ruleRegexp := regexp.MustCompile(`(\d+)\|(\d+)`)
	readRules := true
	accumulator := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readRules = false
		} else if readRules {
			match := ruleRegexp.FindStringSubmatch(line)
			lhs, _ := strconv.Atoi(match[1])
			rhs, _ := strconv.Atoi(match[2])
			rules[lhs] = append(rules[lhs], rhs)
		} else {
			values := collections.Map(slices.Values(strings.Split(line, ",")), maths.AtoiUnwrap)
			var visited []int
			broken := false
			for value := range values {
				newValueRules := rules[value]
				intersection := collections.Intersect(visited, newValueRules)
				if len(intersection) != 0 {
					broken = true
					break
				}
				visited = append(visited, value)
			}
			if !broken {
				collect := slices.Collect(values)
				accumulator += collect[len(collect)/2]
			}
		}
	}
	return strconv.Itoa(accumulator)
}

func part2(reader io.Reader) string {
	rules := make(map[int][]int)
	scanner := bufio.NewScanner(reader)
	ruleRegexp := regexp.MustCompile(`(\d+)\|(\d+)`)
	readRules := true
	accumulator := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readRules = false
		} else if readRules {
			match := ruleRegexp.FindStringSubmatch(line)
			lhs, _ := strconv.Atoi(match[1])
			rhs, _ := strconv.Atoi(match[2])
			rules[lhs] = append(rules[lhs], rhs)
		} else {
			values := collections.Map(slices.Values(strings.Split(line, ",")), maths.AtoiUnwrap)
			var visited []int
			broken := false
			for value := range values {
				newValueRules := rules[value]
				intersection := collections.Intersect(visited, newValueRules)
				if len(intersection) != 0 {
					broken = true
					break
				}
				visited = append(visited, value)
			}
			if broken {
				sorted := slices.SortedFunc(values, func(left int, right int) int {
					leftRules := rules[left]
					rightRules := rules[right]
					if slices.Contains(leftRules, right) {
						return 1
					}
					if slices.Contains(rightRules, left) {
						return -1
					}
					return 0
				})
				accumulator += sorted[len(sorted)/2]
			}

		}
	}
	return strconv.Itoa(accumulator)
}
