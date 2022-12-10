package day02

import (
	"strings"

	"github.com/kaugesaar/advent-of-code/common"
	"github.com/kaugesaar/advent-of-code/utils"
)

// Solution solutions for day 1
type Solution struct{}

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

// Run1 runs the part 1 solution for day 2
func (s Solution) Run1() common.Response {
	return common.Response{
		Output: utils.ToStr(part1(utils.ReadFile("./2022/day02/day2.txt"))),
		Day:    "Day 2",
		Part:   "Part 1",
	}
}

// Run2 runs the part 2 solution for day 2
func (s Solution) Run2() common.Response {
	return common.Response{
		Output: utils.ToStr(part2(utils.ReadFile("./2022/day02/day2.txt"))),
		Day:    "Day 2",
		Part:   "Part 2",
	}
}
