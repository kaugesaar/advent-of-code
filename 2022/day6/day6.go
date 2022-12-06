package main

import (
	_ "embed"
	"fmt"
)

//go:embed day6.txt
var fileInput string

func parser(input string, numChar int) int {
	for i := numChar; i <= len(input); i++ {
		hashset := map[rune]struct{}{}
		for _, c := range input[i-numChar : i] {
			hashset[c] = struct{}{}
		}
		if len(hashset) == numChar {
			return i
		}
	}
	return 0
}

func part1(input string) int {
	return parser(input, 4)
}

func part2(input string) int {
	return parser(input, 14)
}

func main() {
	fmt.Println("---- 2022 Day 6 ----")
	fmt.Println("part1: ", part1(fileInput))
	fmt.Println("part2: ", part2(fileInput))
}
