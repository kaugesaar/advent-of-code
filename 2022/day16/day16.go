package day16

import (
	"regexp"
	"strings"

	"github.com/kaugesaar/advent-of-code/common"
	"github.com/kaugesaar/advent-of-code/utils"
)

// Solution solutions for day 16
type Solution struct{}

type valve struct {
	flowRate   int
	mask       int
	neighbours []string
}

func part1(input string) string {
	valves := parser(input)
	distances := floydWarshall(valves)
	valves = filterVales(valves)
	run := visit("AA", valves, distances, 30, 0, 0, make(map[int]int))
	total := 0
	for _, value := range run {
		if value > total {
			total = value
		}
	}
	return utils.ToStr(total)
}

func part2(input string) string {
	valves := parser(input)
	distances := floydWarshall(valves)
	valves = filterVales(valves)
	run := visit("AA", valves, distances, 26, 0, 0, make(map[int]int))
	total := 0
	for playerState, playerValue := range run {
		for elephantState, elephantValue := range run {
			if (playerState & elephantState) == 0 {
				if playerValue+elephantValue > total {
					total = playerValue + elephantValue
				}
			}
		}
	}
	return utils.ToStr(total)
}

func parser(input string) map[string]valve {
	re := regexp.MustCompile(`Valve ([A-Z]+) has flow rate=([0-9]+); tunnels? leads? to valves? (.*)`)
	valves := make(map[string]valve)
	i := 0
	for _, line := range strings.Split(input, "\n") {
		matches := re.FindStringSubmatch(line)
		name := matches[1]
		flowRate := utils.ToInt(matches[2])
		neighbours := strings.Split(matches[3], ", ")
		valves[name] = valve{
			flowRate:   flowRate,
			mask:       1 << i,
			neighbours: neighbours,
		}
		i++
	}
	return valves
}

func floydWarshall(valves map[string]valve) map[string]int {
	distances := make(map[string]int)
	for from := range valves {
		for to := range valves {
			if utils.Contains(valves[from].neighbours, to) {
				distances[makeKey(from, to)] = 1
			} else {
				distances[makeKey(from, to)] = 99999
			}
		}
	}
	for k := range valves {
		for i := range valves {
			for j := range valves {
				ij := distances[makeKey(i, j)]
				ik := distances[makeKey(i, k)]
				kj := distances[makeKey(k, j)]
				distances[makeKey(i, j)] = utils.MinInt(ij, ik+kj)
			}
		}
	}
	return distances
}

func filterVales(valves map[string]valve) map[string]valve {
	filteredValves := make(map[string]valve)
	for name, valve := range valves {
		if valve.flowRate > 0 {
			filteredValves[name] = valve
		}
	}
	return filteredValves
}

func visit(curr string, valves map[string]valve, distances map[string]int, timeLeft, state, totalFlow int, answers map[int]int) map[int]int {
	n := 0
	if _, ok := answers[state]; ok {
		n = answers[state]
	}
	answers[state] = utils.MaxInt(n, totalFlow)
	for v := range valves {
		distance := distances[makeKey(curr, v)]
		newTimeLeft := timeLeft - distance - 1
		mask := valves[v].mask
		if (state&mask) != 0 || newTimeLeft < 0 {
			continue
		} else {
			flowHere := valves[v].flowRate
			newTotalFlow := totalFlow + (newTimeLeft * flowHere)
			visit(v, valves, distances, newTimeLeft, state|mask, newTotalFlow, answers)
		}
	}
	return answers
}

// Returns key for looking up distance from valve X to valve Y
func makeKey(x, y string) string {
	return x + "," + y
}

// Run1 runs the part 1 solution for day 16
func (s Solution) Run1() common.Response {
	return common.Response{
		Output: part1(utils.ReadFile("./2022/day16/day16.txt")),
		Day:    "Day 16",
		Part:   "Part 1",
	}
}

// Run2 runs the part 2 solution for day 16
func (s Solution) Run2() common.Response {
	return common.Response{
		Output: part2(utils.ReadFile("./2022/day16/day16.txt")),
		Day:    "Day 16",
		Part:   "Part 2",
	}
}
