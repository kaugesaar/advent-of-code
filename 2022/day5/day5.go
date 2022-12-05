package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed day5.txt
var fileInput string

func toInt(s string) int {
	str, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return str
}

func parseCrates(input string) [][]string {
	lines := strings.Split(input, "\n")
	numCols := (len(lines[0]) + 1) / 4
	crates := make([][]string, numCols)
	for j := len(lines) - 1; j > -1; j-- {
		for i, crate := range crates {
			c := lines[j][1+4*i]
			if c >= 'A' && c <= 'Z' {
				crates[i] = append(crate, string(c))
			}
		}
	}
	return crates
}

func parseMoves(input string) [][]int {
	lines := strings.Split(input, "\n")
	var moves [][]int
	re := regexp.MustCompile("([0-9]{1,999})")
	for _, line := range lines {
		match := re.FindAllString(line, 3)
		move := []int{toInt(match[0]), toInt(match[1]), toInt(match[2])}
		moves = append(moves, move)
	}
	return moves
}

func part2(input string) string {
	parts := strings.Split(string(input), "\n\n")
	crates := parseCrates(parts[0])
	moves := parseMoves(parts[1])
	for _, move := range moves {
		m := move[0]
		f := move[1]
		t := move[2]
		split := len(crates[f-1]) - m
		top := crates[f-1][split:]
		crates[f-1] = crates[f-1][:split]
		crates[t-1] = append(crates[t-1], top...)
	}
	var result string
	for _, crate := range crates {
		lastCrate := crate[len(crate)-1]
		result += string(lastCrate)
	}
	return result
}

func part1(input string) string {
	parts := strings.Split(string(input), "\n\n")
	crates := parseCrates(parts[0])
	moves := parseMoves(parts[1])
	for _, move := range moves {
		m := move[0]
		f := move[1]
		t := move[2]
		for i := 0; i < m; i++ {
			top := crates[f-1][len(crates[f-1])-1]
			crates[f-1] = crates[f-1][:len(crates[f-1])-1]
			crates[t-1] = append(crates[t-1], top)
		}
	}
	var result string
	for _, crate := range crates {
		lastCrate := crate[len(crate)-1]
		result += string(lastCrate)
	}
	return result
}

func main() {
	fmt.Println("---- 2022 Day 5 ----")
	fmt.Println("part1: ", part1(fileInput))
	fmt.Println("part2: ", part2(fileInput))
}
