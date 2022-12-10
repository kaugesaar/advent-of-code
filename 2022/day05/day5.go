package day05

import (
	"regexp"
	"strings"

	"github.com/kaugesaar/advent-of-code/common"
	"github.com/kaugesaar/advent-of-code/utils"
)

// Solution solutions for day 5
type Solution struct{}

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
		move := []int{utils.ToInt(match[0]), utils.ToInt(match[1]), utils.ToInt(match[2])}
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

// Run1 runs the part 1 solution for day 1
func (s Solution) Run1() common.Response {
	return common.Response{
		Output: part1(utils.ReadFile("./2022/day05/day5.txt")),
		Day:    "Day 1",
		Part:   "Part 1",
	}
}

// Run2 runs the part 2 solution for day 1
func (s Solution) Run2() common.Response {
	return common.Response{
		Output: part2(utils.ReadFile("./2022/day05/day5.txt")),
		Day:    "Day 1",
		Part:   "Part 2",
	}
}
