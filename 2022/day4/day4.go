package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day4.txt
var fileInput string

func fullyContains(a []int, b []int) bool {
	return (a[0] <= b[0] && a[1] >= b[1]) || (b[0] <= a[0] && b[1] >= a[1])
}

func overlaps(a []int, b []int) bool {
	return (a[0] <= b[0] && b[0] <= a[1]) || (b[0] <= a[0] && a[0] <= b[1])
}

func parseSection(assignments string) []int {
	sections := strings.Split(assignments, "-")
	low, _ := strconv.Atoi(sections[0])
	high, _ := strconv.Atoi(sections[1])
	return []int{low, high}
}

func parsePair(pairs string) ([]int, []int) {
	assignments := strings.Split(pairs, ",")
	a := parseSection(assignments[0])
	b := parseSection(assignments[1])
	return a, b
}

func parser(lines []string, checker func([]int, []int) bool) int {
	var count int
	for _, pairs := range lines {
		a, b := parsePair(pairs)
		if checker(a, b) {
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
