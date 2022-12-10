package day06

import (
	"github.com/kaugesaar/advent-of-code/common"
	"github.com/kaugesaar/advent-of-code/utils"
)

// Solution solutions for day 6
type Solution struct{}

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

// Run1 runs the part 1 solution for day 6
func (s Solution) Run1() common.Response {
	return common.Response{
		Output: utils.ToStr(part1(utils.ReadFile("./2022/day06/day6.txt"))),
		Day:    "Day 6",
		Part:   "Part 1",
	}
}

// Run2 runs the part 2 solution for day 6
func (s Solution) Run2() common.Response {
	return common.Response{
		Output: utils.ToStr(part2(utils.ReadFile("./2022/day06/day6.txt"))),
		Day:    "Day 6",
		Part:   "Part 2",
	}
}
