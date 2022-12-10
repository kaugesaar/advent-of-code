package day01

import (
	"sort"
	"strings"

	"github.com/kaugesaar/advent-of-code/common"
	"github.com/kaugesaar/advent-of-code/utils"
)

// Solution solutions for day 1
type Solution struct{}

func parser(parts []string) []int {
	var sums []int
	for _, part := range parts {
		var sum int
		for _, line := range strings.Split(part, "\n") {
			sum += utils.ToInt(line)
		}
		sums = append(sums, sum)
	}
	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})
	return sums
}

func part1(input string) int {
	lines := strings.Split(input, "\n\n")
	result := parser(lines)
	return result[0]
}

func part2(input string) int {
	lines := strings.Split(input, "\n\n")
	result := parser(lines)
	return result[0] + result[1] + result[2]
}

// Run1 runs the part 1 solution for day 1
func (s Solution) Run1() common.Response {
	return common.Response{
		Output: utils.ToStr(part1(utils.ReadFile("./2022/day01/day1.txt"))),
		Day:    "Day 1",
		Part:   "Part 1",
	}
}

// Run2 runs the part 2 solution for day 1
func (s Solution) Run2() common.Response {
	return common.Response{
		Output: utils.ToStr(part2(utils.ReadFile("./2022/day01/day1.txt"))),
		Day:    "Day 1",
		Part:   "Part 2",
	}
}
