package day12

import (
	"strings"

	"github.com/kaugesaar/advent-of-code/common"
	"github.com/kaugesaar/advent-of-code/utils"
	"github.com/kaugesaar/advent-of-code/utils/matrix"
	"github.com/kaugesaar/advent-of-code/utils/pathfinding"
)

// Solution solutions for day 12
type Solution struct{}

func part1(input string) string {
	rows := strings.Split(input, "\n")
	m := matrix.BuildIntMatrix(rows)
	var start matrix.Pos
	var end matrix.Pos

	for j, l := range m {
		for i, c := range l {
			if c == 'S' {
				start = matrix.Pos{X: i, Y: j}
			}
			if c == 'E' {
				end = matrix.Pos{X: i, Y: j}
			}
		}
	}

	m[start.Y][start.X] = 'a'
	m[end.Y][end.X] = 'z'

	return utils.ToStr(pathfinding.Bfs(m, start, end))
}

func part2(input string) string {
	rows := strings.Split(input, "\n")
	m := matrix.BuildIntMatrix(rows)
	var start matrix.Pos
	best := 0

	for y, l := range m {
		for x, c := range l {
			if c == 'S' {
				m[y][x] = 'a'
			}
			if c == 'E' {
				start = matrix.Pos{X: x, Y: y}
				m[y][x] = 'z'
			}
		}
	}

	best = pathfinding.BfsFindValue(m, start, 'a')

	return utils.ToStr(best)
}

// Run1 runs the part 1 solution for day 12
func (s Solution) Run1() common.Response {
	return common.Response{
		Output: part1(utils.ReadFile("./2022/day12/day12.txt")),
		Day:    "Day 12",
		Part:   "Part 1",
	}
}

// Run2 runs the part 2 solution for day 12
func (s Solution) Run2() common.Response {
	return common.Response{
		Output: part2(utils.ReadFile("./2022/day12/day12.txt")),
		Day:    "Day 12",
		Part:   "Part 2",
	}
}
