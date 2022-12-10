package day07

import (
	"strings"

	"github.com/kaugesaar/advent-of-code/common"
	"github.com/kaugesaar/advent-of-code/utils"
)

// Solution solutions for day 7
type Solution struct{}

func parser(lines []string, s func([]int) int) int {
	var dir []int
	var directories []int

	for _, line := range lines {
		if strings.HasPrefix(line, "$ cd") {
			if strings.Contains(line, "..") {
				directories = append(directories, dir[len(dir)-1])
				dir = dir[:len(dir)-1]
			} else {
				dir = append(dir, 0)
			}
		}
		if !strings.HasPrefix(line, "$") && !strings.HasPrefix(line, "dir") {
			fields := strings.Fields(line)
			size := utils.ToInt(fields[0])
			for i := range dir {
				dir[i] += size
			}
		}
	}
	directories = append(directories, dir...)
	return s(directories)
}

func solution1(directories []int) int {
	var sum int
	for _, size := range directories {
		if size <= 100000 {
			sum += size
		}
	}
	return sum
}

func solution2(directories []int) int {
	max := 0
	for _, size := range directories {
		if size > max {
			max = size
		}
	}
	req := 30000000 - (70000000 - max)
	min := 0
	for _, size := range directories {
		if size >= req && (min == 0 || size < min) {
			min = size
		}
	}
	return min
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	return parser(lines, solution1)
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	return parser(lines, solution2)
}

// Run1 runs the part 1 solution for day 7
func (s Solution) Run1() common.Response {
	return common.Response{
		Output: utils.ToStr(part1(utils.ReadFile("./2022/day07/day7.txt"))),
		Day:    "Day 7",
		Part:   "Part 1",
	}
}

// Run2 runs the part 2 solution for day 7
func (s Solution) Run2() common.Response {
	return common.Response{
		Output: utils.ToStr(part2(utils.ReadFile("./2022/day07/day7.txt"))),
		Day:    "Day 7",
		Part:   "Part 2",
	}
}
