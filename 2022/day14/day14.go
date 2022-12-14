package day14

import (
	"strings"

	"github.com/kaugesaar/advent-of-code/common"
	"github.com/kaugesaar/advent-of-code/utils"
)

// Solution solutions for day 14
type Solution struct{}

type point struct {
	x int
	y int
}

func part1(input string) string {
	grid := parse(input)
	count := 0
	var maxY int
	for p := range grid {
		maxY = utils.MaxInt(maxY, p.y)
	}
	var sand point
	for sand.y <= maxY {
		sand = point{x:500, y:0}
		for !grid[sand] && sand.y <= maxY {
			if !grid[point{sand.x, sand.y + 1}] {
				sand.y += 1
			} else if !grid[point{sand.x - 1, sand.y + 1}] {
				sand.x -= 1
				sand.y += 1
			} else if !grid[point{sand.x + 1, sand.y + 1}] {
				sand.x += 1
				sand.y += 1
			} else {
				grid[sand] = true
				count++
			}
		}
	}
	return utils.ToStr(count)
}

func part2(input string) string {
	grid := parse(input)
	count := 0
	var maxY int
	for p := range grid {
		maxY = utils.MaxInt(maxY, p.y)
	}
	floor := maxY + 2
	var sand point
	for !grid[point{500, 0}] {
		sand = point{x:500, y:0}
		for !grid[sand] && sand.y <= floor - 1 {
			if !grid[point{sand.x, sand.y + 1}] {
				sand.y += 1
			} else if !grid[point{sand.x - 1, sand.y + 1}] {
				sand.x -= 1
				sand.y += 1
			} else if !grid[point{sand.x + 1, sand.y + 1}] {
				sand.x += 1
				sand.y += 1
			} else {
				grid[sand] = true
				count++
			}
			if sand.y + 1 == floor {
				grid[sand] = true
				count++
			}
		}
	}
	return utils.ToStr(count)
}

func parse(input string) map[point]bool {
	grid := map[point]bool{}
	
	for _, line := range strings.Split(input, "\n") {
		var prevPoint point

		for i, p := range strings.Split(line, " -> ") {
			x, y := utils.ToInt(strings.Split(p, ",")[0]), utils.ToInt(strings.Split(p, ",")[1])
			currPoint := point{x:x, y:y}

			if i > 0 {
				for i := utils.MinInt(currPoint.x, prevPoint.x); i <= utils.MaxInt(currPoint.x, prevPoint.x); i++ {
					grid[point{x:i, y:y}] = true
				}
				for j := utils.MinInt(currPoint.y, prevPoint.y); j <= utils.MaxInt(currPoint.y, prevPoint.y); j++ {
					grid[point{x:x, y:j}] = true
				}
			} else {
				grid[currPoint] = true
			}
			prevPoint = currPoint
		}
	}
	return grid
}

// Run1 runs the part 1 solution for day 14
func (s Solution) Run1() common.Response {
	return common.Response{
		Output: part1(utils.ReadFile("./2022/day14/day14.txt")),
		Day:    "Day 14",
		Part:   "Part 1",
	}
}

// Run2 runs the part 2 solution for day 14
func (s Solution) Run2() common.Response {
	return common.Response{
		Output: part2(utils.ReadFile("./2022/day14/day14.txt")),
		Day:    "Day 14",
		Part:   "Part 2",
	}
}
