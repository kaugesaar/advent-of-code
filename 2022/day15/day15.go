package day15

import (
	"fmt"
	"strings"

	"github.com/kaugesaar/advent-of-code/common"
	"github.com/kaugesaar/advent-of-code/utils"
)

// Solution solutions for day 15
type Solution struct{}

type point struct{ x, y int }

func (p *point) equal(b point) bool {
	return p.x == b.x && p.y == b.y
}

type sensor struct {
	pos      point
	beacon   point
	distance int
}

func (s *sensor) contains(b point) bool {
	return s.distance >= manhattan(s.pos, b)
}

func part1(input string, row int) string {
	sensors := parser(input)
	xMin, xMax := xMinxMax(sensors)
	count := 0

	for x := xMin; x <= xMax; x++ {
		testPoint := point{x, row}
		for _, sensor := range sensors {
			if sensor.contains(testPoint) &&
				!sensor.pos.equal(testPoint) &&
				!sensor.beacon.equal(testPoint) {
				count++
				break
			}
		}
	}

	return utils.ToStr(count)
}

func part2(input string, coordMax int) string {
	sensors := parser(input)
	possibleBeacon := point{}
	coveringSensor := sensor{}

	for {
		covered := false
		for _, sensor := range sensors {
			covered = sensor.contains(possibleBeacon)
			if covered {
				coveringSensor = sensor
				break
			}
		}

		if !covered {
			break
		}

		skip := coveringSensor.distance - manhattan(coveringSensor.pos, possibleBeacon) + 1

		if possibleBeacon.x+skip > coordMax {
			possibleBeacon.x = 0
			possibleBeacon.y++
		} else {
			possibleBeacon.x += skip
		}

	}
	return utils.ToStr(possibleBeacon.x*4000000 + possibleBeacon.y)
}

func manhattan(a point, b point) int {
	return utils.Abs((a.x - b.x)) + utils.Abs((a.y - b.y))
}

func xMinxMax(sensors []sensor) (int, int) {
	xMin := 99999999
	xMax := 0

	for _, sensor := range sensors {
		if sensor.pos.x-sensor.distance < xMin {
			xMin = sensor.pos.x - sensor.distance
		}
		if sensor.pos.x+sensor.distance > xMax {
			xMax = sensor.pos.x + sensor.distance
		}
	}

	return xMin, xMax
}

func parser(input string) []sensor {
	var sensors []sensor
	for _, line := range strings.Split(input, "\n") {
		var sX, sY, bX, bY int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sX, &sY, &bX, &bY)
		p := point{sX, sY}
		b := point{bX, bY}
		sensors = append(sensors, sensor{
			pos:      p,
			beacon:   b,
			distance: manhattan(p, b),
		})
	}
	return sensors
}

// Run1 runs the part 1 solution for day 15
func (s Solution) Run1() common.Response {
	return common.Response{
		Output: part1(utils.ReadFile("./2022/day15/day15.txt"), 2000000),
		Day:    "Day 15",
		Part:   "Part 1",
	}
}

// Run2 runs the part 2 solution for day 15
func (s Solution) Run2() common.Response {
	return common.Response{
		Output: part2(utils.ReadFile("./2022/day15/day15.txt"), 4000000),
		Day:    "Day 15",
		Part:   "Part 2",
	}
}
