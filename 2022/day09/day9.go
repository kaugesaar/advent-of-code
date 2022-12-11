package day09

import (
	"math"
	"strconv"
	"strings"

	"github.com/kaugesaar/advent-of-code/common"
	"github.com/kaugesaar/advent-of-code/utils"
)

// Solution solutions for day 9
type Solution struct{}

type Pos struct {
	x int32
	y int32
}

func (p *Pos) Init(x int32, y int32) *Pos {
	p.x = x
	p.y = y
	return p
}

func (p *Pos) Signum() *Pos {
	if p.x < 0 {
		p.x = -1
	}
	if p.x > 0 {
		p.x = 1
	}
	if p.y < 0 {
		p.y = -1
	}
	if p.y > 0 {
		p.y = 1
	}
	return p
}

func (p *Pos) Abs() *Pos {
	p.x = int32(math.Abs(float64(p.x)))
	p.y = int32(math.Abs(float64(p.y)))
	return p
}

func (p *Pos) Max() int32 {
	return int32(math.Max(float64(p.x), float64(p.y)))
}

func (p *Pos) ParseInput(imp string) (Pos, int32) {
	parts := strings.SplitN(imp, " ", 2)
	dir := parts[0]
	count := parts[1]
	countInt, _ := strconv.Atoi(count)
	var pos Pos
	switch dir {
	case "R":
		pos = *p.Init(1, 0)
	case "D":
		pos = *p.Init(0, 1)
	case "L":
		pos = *p.Init(-1, 0)
	case "U":
		pos = *p.Init(0, -1)
	default:
		panic("Movement parsing error")
	}
	return pos, int32(countInt)
}

func parser(input string, n int) string {
	var pos = Pos{}
	movements := strings.Split(input, "\n")
	var tailMap = make(map[Pos]struct{})
	knots := make([]Pos, n)
	for _, m := range movements {
		dir, count := pos.ParseInput(m)
		for i := 0; i < int(count); i++ {
			knots[0].y += dir.y
			knots[0].x += dir.x
			for j := 1; j < n; j++ {
				diff := Pos{
					x: knots[j-1].x - knots[j].x,
					y: knots[j-1].y - knots[j].y,
				}
				if diff.Abs().Max() <= 1 {
					continue
				}
				d := Pos{
					x: knots[j-1].x - knots[j].x,
					y: knots[j-1].y - knots[j].y,
				}
				signum := d.Signum()
				knots[j].x += signum.x
				knots[j].y += signum.y

			}
			tailMap[knots[len(knots)-1]] = struct{}{}
		}
	}
	return strconv.Itoa(len(tailMap))
}

func part1(input string) string {
	return parser(input, 2)
}

func part2(input string) string {
	return parser(input, 10)
}

// Run1 runs the part 1 solution for day 9
func (s Solution) Run1() common.Response {
	return common.Response{
		Output: part1(utils.ReadFile("./2022/day09/day9.txt")),
		Day:    "Day 9",
		Part:   "Part 1",
	}
}

// Run2 runs the part 2 solution for day 9
func (s Solution) Run2() common.Response {
	return common.Response{
		Output: part2(utils.ReadFile("./2022/day09/day9.txt")),
		Day:    "Day 9",
		Part:   "Part 2",
	}
}
