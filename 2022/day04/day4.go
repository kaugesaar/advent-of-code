package day04

import (
	"fmt"
	"strings"

	"github.com/kaugesaar/advent-of-code/common"
	"github.com/kaugesaar/advent-of-code/utils"
)

// Solution solutions for day 4
type Solution struct{}

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

// Run1 runs the part 1 solution for day 4
func (s Solution) Run1() common.Response {
	return common.Response{
		Output: utils.ToStr(part1(utils.ReadFile("./2022/day04/day4.txt"))),
		Day:    "Day 4",
		Part:   "Part 1",
	}
}

// Run2 runs the part 2 solution for day 4
func (s Solution) Run2() common.Response {
	return common.Response{
		Output: utils.ToStr(part2(utils.ReadFile("./2022/day04/day4.txt"))),
		Day:    "Day 4",
		Part:   "Part 2",
	}
}
