package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed day3.txt
var fileInput string

func score(item rune) int {
	switch {
	case 'A' <= item && item <= 'Z':
		return int(item-'A') + 27
	case 'a' <= item && item <= 'z':
		return int(item-'a') + 1
	default:
		return 0
	}
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	var sum int
	for i := 0; i < len(lines); i += 3 {
		r1 := lines[i]
		r2 := lines[i+1]
		r3 := lines[i+2]
		for _, c := range r1 {
			if strings.ContainsRune(r2, c) && strings.ContainsRune(r3, c) {
				sum += score(c)
				break
			}
		}
	}
	return sum
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	var sum int
	for _, rucksack := range lines {
		a := rucksack[:len(rucksack)/2]
		b := rucksack[len(rucksack)/2:]
		for _, c := range a {
			if strings.ContainsRune(b, c) {
				sum += score(c)
				break
			}
		}
	}
	return sum
}

func main() {
	fmt.Println("---- 2022 Day 3 ----")
	fmt.Println("part1: ", part1(fileInput))
	fmt.Println("part2: ", part2(fileInput))
}
