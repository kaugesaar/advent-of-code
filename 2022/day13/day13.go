package day13

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/kaugesaar/advent-of-code/common"
	"github.com/kaugesaar/advent-of-code/utils"
)

// Solution solutions for day 13
type Solution struct{}

type output struct {
	values []int
}

func part1(input string) string {
	packets := parseInput(input)
	result := 0
	for i, pair := range packets {
		if comparePair(pair[0], pair[1]) <= 0 {
			result += i + 1
		}
	}
	return utils.ToStr(result)
}

func part2(input string) string {
	parsedInput := parseInput(input)
	packets := []any{}
	for _, pair := range parsedInput {
		packets = append(packets, pair...)
	}

	var divider1 any
	json.Unmarshal([]byte("[[2]]"), &divider1)

	var divider2 any
	json.Unmarshal([]byte("[[6]]"), &divider2)

	packets = append(packets, []any{divider1, divider2}...)
	sort.Slice(packets, func(i, j int) bool {
		return comparePair(packets[i], packets[j]) <= 0
	})

	result := 1
	for i, packet := range packets {
		packet, _ := json.Marshal(packet)
		s := string(packet)
		if s == "[[2]]" || s == "[[6]]" {
			result *= i + 1
		}
	}
	return utils.ToStr(result)
}

func comparePair(p1, p2 any) int {
	_, ok1 := p1.(float64)
	_, ok2 := p2.(float64)
	if ok1 && ok2 {
		return int(p1.(float64) - p2.(float64))
	}
	if ok1 {
		p1 = []any{p1}
	}
	if ok2 {
		p2 = []any{p2}
	}
	if len(p1.([]any)) == 0 || len(p2.([]any)) == 0 {
		return len(p1.([]any)) - len(p2.([]any))
	}
	result := comparePair(p1.([]any)[0], p2.([]any)[0])
	if result == 0 {
		next1 := p1.([]any)[1:]
		next2 := p2.([]any)[1:]
		if len(next1) == 0 || len(next2) == 0 {
			return len(next1) - len(next2)
		}
		return comparePair(next1, next2)
	}
	return result
}

func parseInput(input string) [][]any {
	lines := strings.Split(input, "\n")
	packets := [][]any{}
	pair := []any{}

	for _, line := range lines {
		if line == "" {
			packets = append(packets, pair)
			pair = []any{}
		} else {
			var p any
			err := json.Unmarshal([]byte(line), &p)
			if err != nil {
				panic("Error with unmarshal")
			}
			pair = append(pair, p)
		}
	}

	return append(packets, pair)
}

// Run1 runs the part 1 solution for day 13
func (s Solution) Run1() common.Response {
	return common.Response{
		Output: part1(utils.ReadFile("./2022/day13/day13.txt")),
		Day:    "Day 13",
		Part:   "Part 1",
	}
}

// Run2 runs the part 2 solution for day 13
func (s Solution) Run2() common.Response {
	return common.Response{
		Output: part2(utils.ReadFile("./2022/day13/day13.txt")),
		Day:    "Day 13",
		Part:   "Part 2",
	}
}
