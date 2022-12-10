package day03

import (
	"strings"

	"github.com/kaugesaar/advent-of-code/common"
	"github.com/kaugesaar/advent-of-code/utils"
)

// Solution solutions for day 3
type Solution struct{}

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

// Run1 runs the part 1 solution for day 3
func (s Solution) Run1() common.Response {
	return common.Response{
		Output: utils.ToStr(part1(utils.ReadFile("./2022/day03/day3.txt"))),
		Day:    "Day 3",
		Part:   "Part 1",
	}
}

// Run2 runs the part 2 solution for day 3
func (s Solution) Run2() common.Response {
	return common.Response{
		Output: utils.ToStr(part2(utils.ReadFile("./2022/day03/day3.txt"))),
		Day:    "Day 3",
		Part:   "Part 2",
	}
}
