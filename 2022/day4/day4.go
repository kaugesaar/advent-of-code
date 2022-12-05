package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed day4.txt
var fileInput string

func fullyContains(a int, b int, c int, d int) bool {
	return (a <= c && b >= d) || (c <= a && d >= b)
}

func overlaps(a int, b int, c int, d int) bool {
	return !(b < c || d < a)
}

func parser(lines []string, checker func(int, int, int, int) bool) int {
	var count int
	for _, line := range lines {
		var a, b, c, d int
		fmt.Sscanf(line, "%d-%d,%d-%d", &a, &b, &c, &d)
		if checker(a, b, c, d) {
			count++
		}
	}
	return count
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	result := parser(lines, fullyContains)
	return result
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	result := parser(lines, overlaps)
	return result
}

func main() {
	fmt.Println("---- 2022 Day 4 ----")
	fmt.Println("part1: ", part1(fileInput))
	fmt.Println("part2: ", part2(fileInput))
}
