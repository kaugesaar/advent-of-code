package day10

import (
	"strconv"
	"strings"

	"github.com/kaugesaar/advent-of-code/common"
	"github.com/kaugesaar/advent-of-code/utils"
)

// Solution solutions for day 10
type Solution struct{}

type row struct {
	process string
	cycle   int
}

func part1(input string) string {
	return parser(input, solution1)
}

func part2(input string) string {
	return parser(input, solution2)
}

func contains(s []int, i int) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}

func solution1(lines []row) string {
	cycle := 1
	cycles := []int{20, 60, 100, 140, 180, 220}
	x := 1
	signalStrength := 0

	for _, line := range lines {
		iterations := 1
		if line.process == "addx" {
			iterations = 2
		}
		for i := 0; i < iterations; i++ {
			if contains(cycles, cycle) {
				signalStrength += x * cycle
			}
			cycle++
		}
		x += line.cycle
	}
	return strconv.Itoa(signalStrength)
}

func solution2(lines []row) string {
	x := 1
	rows := make([]string, 6)
	col := 0
	row := 0
	for _, line := range lines {
		iterations := 1
		if line.process == "addx" {
			iterations = 2
		}
		for i := 0; i < iterations; i++ {
			pixel := "."
			if col == x || col == x-1 || col == x+1 {
				pixel = "#"
			}
			rows[row] += pixel
			col++
			if col == 40 {
				col = 0
				row++
			}
		}
		x += line.cycle
	}

	result := "\n"
	for _, r := range rows {
		result += r + "\n"
	}
	return result
}

func parser(input string, f func([]row) string) string {
	var result = []row{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		r := strings.Split(line, " ")
		row := row{r[0], 0}
		if len(r) == 2 {
			row.cycle = utils.ToInt(r[1])
		}

		result = append(result, row)
	}
	return f(result)
}

// Run1 runs the part 1 solution for day 10
func (s Solution) Run1() common.Response {
	return common.Response{
		Output: part1(utils.ReadFile("./2022/day10/day10.txt")),
		Day:    "Day 10",
		Part:   "Part 1",
	}
}

// Run2 runs the part 2 solution for day 10
func (s Solution) Run2() common.Response {
	return common.Response{
		Output: part2(utils.ReadFile("./2022/day10/day10.txt")),
		Day:    "Day 10",
		Part:   "Part 2",
	}
}
