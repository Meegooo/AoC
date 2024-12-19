package main

import (
	"bufio"
	"github.com/meegoue/AoC/library/collections"
	"github.com/meegoue/AoC/library/maths"
	"io"
	"iter"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day2/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	part2(file)
}

func part1(reader io.Reader) string {
	safe := 0
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		splitted := strings.Split(text, " ")
		ints := collections.Map(slices.Values(splitted), func(t string) int {
			res, err := strconv.Atoi(t)
			if err != nil {
				log.Fatal(err)
			}
			return res
		})

		if check(ints) {
			safe++
		}
	}
	return strconv.Itoa(safe)

}

func part2(reader io.Reader) string {
	safe := 0
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		splitted := strings.Split(text, " ")
		ints := collections.Map(slices.Values(splitted), func(t string) int {
			res, err := strconv.Atoi(t)
			if err != nil {
				log.Fatal(err)
			}
			return res
		})

		if check(ints) {
			safe++
		} else {
			skipIdx := 0
			for range ints {
				if check(collections.SkipOne(ints, skipIdx)) {
					safe++
					break
				}
				skipIdx++
			}
		}
	}
	return strconv.Itoa(safe)
}

func check(ints iter.Seq[int]) bool {
	increasing := false
	prev := -1
	idx := 0
	skip := false
	for i := range ints {
		if idx > 0 {
			diff := maths.Abs(i - prev)
			if idx == 1 {
				if i < prev {
					increasing = false
				} else if i > prev {
					increasing = true
				}
			}
			if diff < 1 || diff > 3 || (increasing && i < prev) || (!increasing && i > prev) {
				skip = true
				break
			}
		}
		prev = i
		idx++
	}

	return !skip
}
