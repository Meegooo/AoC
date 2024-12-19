package main

import (
	"bufio"
	"fmt"
	"github.com/meegoue/AoC/library/collections"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day9/test.txt")
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
	input := slices.Collect(collections.Map(slices.Values([]rune(scanner.Text())), func(t rune) int {
		return int(t - '0')
	}))
	checksum := int64(0)
	rightDigitId := len(input)/2 + 1

	var diskPosition int64 = 0

	rightDigitRemaining := 0

	result := make([]rune, 0)

	for leftPtr, digit := range input {
		if rightDigitId*2 < leftPtr {
			break
		}
		if leftPtr%2 == 0 { // file
			if leftPtr == rightDigitId*2 {
				for range rightDigitRemaining {
					checksum += int64(leftPtr/2) * diskPosition
					result = append(result, rune(leftPtr/2+'0'))
					diskPosition++
				}

			} else {
				for range digit {
					checksum += int64(leftPtr/2) * diskPosition
					result = append(result, rune(leftPtr/2+'0'))
					diskPosition++
				}
			}
		} else if rightDigitId*2 > leftPtr {
			for range digit {
				for rightDigitRemaining == 0 {
					rightDigitId--
					if rightDigitId*2 <= leftPtr {
						break
					}
					rightDigitRemaining = input[rightDigitId*2]
				}
				if rightDigitId*2 <= leftPtr {
					break
				}
				checksum += int64(rightDigitId) * diskPosition
				result = append(result, rune(rightDigitId+'0'))
				diskPosition++
				rightDigitRemaining--
			}
		}
	}
	//fmt.Println(string(result))
	return strconv.FormatInt(checksum, 10)
}

type entry struct {
	size  int
	id    int
	space bool
}

func part2(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	var input []entry
	for i, r := range scanner.Text() {
		input = append(input, entry{size: int(r - '0'), id: i / 2, space: i%2 == 1})
	}

	for rightIdx := len(input) - 1; rightIdx >= 0; rightIdx-- {
		elem := input[rightIdx]
		if elem.space {
			continue
		}
		for leftIdx := 0; leftIdx < rightIdx; leftIdx++ {
			space := input[leftIdx]
			if !space.space || space.size < elem.size {
				continue
			}
			if space.size == elem.size {
				input[leftIdx] = elem
				input[rightIdx].space = true
			} else {
				input[leftIdx] = entry{
					size:  elem.size,
					id:    elem.id,
					space: false,
				}
				newSpace := entry{
					size:  space.size - elem.size,
					id:    space.id,
					space: true,
				}
				input[rightIdx].space = true
				input = slices.Insert(input, leftIdx+1, newSpace)
				rightIdx++
			}
			break
		}

		// Merge spaces on right
		if input[rightIdx].space {
			if rightIdx < len(input)-1 {
				righterElem := input[rightIdx+1]
				if righterElem.space {
					input[rightIdx+1].size = 0
					input[rightIdx].size += righterElem.size
				}
			}
			if rightIdx > 0 {
				lefterElem := input[rightIdx-1]
				if lefterElem.space {
					input[rightIdx-1].size = 0
					input[rightIdx].size += lefterElem.size
				}
			}
		}
	}
	accumulator := int64(0)

	idx := 0
	for _, e := range input {
		if e.space {
			idx += e.size
		} else {
			for range e.size {
				accumulator += int64(e.id * idx)
				idx++
			}
		}
	}
	printInput(input)
	return strconv.FormatInt(accumulator, 10)
}

func printInput(input []entry) {
	p := make([]rune, 0)
	for _, e := range input {
		if e.space {
			for range e.size {
				p = append(p, '.')
			}
		} else {
			for range e.size {
				p = append(p, rune('0'+e.id))
			}
		}
	}
	fmt.Println(string(p))
}
