package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type move struct {
	move int
	from int
	to   int
}

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

func parseMoves(input string) []move {
	lines := strings.Split(input, "\n")

	var moves []move

	re := regexp.MustCompile("([0-9]{1,999})")

	for _, line := range lines {
		match := re.FindAllString(line, 3)
		move := move{move: toInt(match[0]), from: toInt(match[1]), to: toInt(match[2])}
		moves = append(moves, move)
	}

	return moves

}

func questionOne(cratesInput string, movesInput string) string {
	crates := parseCrates(cratesInput)
	moves := parseMoves(movesInput)
	for _, move := range moves {
		m := move.move
		f := move.from
		t := move.to
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

func questionTwo(cratesInput string, movesInput string) string {
	crates := parseCrates(cratesInput)
	moves := parseMoves(movesInput)
	for _, move := range moves {
		m := move.move
		f := move.from
		t := move.to
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

func main() {
	input, err := ioutil.ReadFile("./input/day5.txt")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(input), "\n\n")
	cratesInput := parts[0]
	movesInput := parts[1]

	fmt.Println("Day 2 | Question 1:", questionOne(cratesInput, movesInput))
	fmt.Println("Day 2 | Question 2:", questionTwo(cratesInput, movesInput))

}
