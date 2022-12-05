package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed day2.txt
var fileInput string

var scores = map[string]map[string]int{
	"part1": {
		"A X": 4, "A Y": 8, "A Z": 3,
		"B X": 1, "B Y": 5, "B Z": 9,
		"C X": 7, "C Y": 2, "C Z": 6,
	},
	"part2": {
		"A X": 3, "A Y": 4, "A Z": 8,
		"B X": 1, "B Y": 5, "B Z": 9,
		"C X": 2, "C Y": 6, "C Z": 7,
	},
}

func parser(lines []string, part string) int {
	var sum int
	for _, game := range lines {
		if score, ok := scores[part][game]; ok {
			sum += score
		}
	}
	return sum
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	return parser(lines, "part1")
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	return parser(lines, "part2")
}

func main() {
	fmt.Println("---- 2022 Day 2 ----")
	fmt.Println("part1: ", part1(fileInput))
	fmt.Println("part2: ", part2(fileInput))
}
