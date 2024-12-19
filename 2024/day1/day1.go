package main

import (
	"bufio"
	"fmt"
	"github.com/meegoue/AoC/library/maths"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day1/test.txt")
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
	var leftArr []int
	var rightArr []int
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		splitted := strings.Split(text, "  ")
		leftInt, _ := strconv.Atoi(strings.TrimLeft(splitted[0], " "))
		rightInt, _ := strconv.Atoi(strings.TrimLeft(splitted[1], " "))
		leftArr = append(leftArr, leftInt)
		rightArr = append(rightArr, rightInt)
	}
	fmt.Println(leftArr)
	fmt.Println(rightArr)
	sort.Ints(leftArr)
	sort.Ints(rightArr)
	totalDistance := 0

	for i := 0; i < len(leftArr); i++ {
		totalDistance += maths.Abs(leftArr[i] - rightArr[i])
	}
	return strconv.Itoa(totalDistance)
}

func part2(reader io.Reader) string {
	var leftArr []int
	var rightCounter = make(map[int]int)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		splitted := strings.Split(text, "  ")
		leftInt, _ := strconv.Atoi(strings.TrimLeft(splitted[0], " "))
		rightInt, _ := strconv.Atoi(strings.TrimLeft(splitted[1], " "))
		leftArr = append(leftArr, leftInt)
		rightCounter[rightInt]++
	}

	totalDistance := 0

	for i := 0; i < len(leftArr); i++ {
		totalDistance += leftArr[i] * rightCounter[leftArr[i]]
	}
	return strconv.Itoa(totalDistance)
}
