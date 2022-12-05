package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed day1.txt
var fileInput string

func toInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func parser(parts []string) []int {
	var sums []int
	for _, part := range parts {
		var sum int
		for _, line := range strings.Split(part, "\n") {
			sum += toInt(line)
		}
		sums = append(sums, sum)
	}
	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})
	return sums
}

func part1(input string) int {
	lines := strings.Split(input, "\n\n")
	result := parser(lines)
	return result[0]
}

func part2(input string) int {
	lines := strings.Split(input, "\n\n")
	result := parser(lines)
	return result[0] + result[1] + result[2]
}

func main() {
	fmt.Println("---- 2022 Day 1 ----")
	fmt.Println("part1: ", part1(fileInput))
	fmt.Println("part2: ", part2(fileInput))
}
