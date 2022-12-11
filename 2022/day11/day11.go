package day11

import (
	"regexp"
	"sort"
	"strings"

	"github.com/kaugesaar/advent-of-code/common"
	"github.com/kaugesaar/advent-of-code/utils"
)

// Solution solutions for day 11
type Solution struct{}

type monkey struct {
	items          []int64
	inspectedItems int
	operation      operation
	test           test
}

type operation struct {
	increasedBy  int64
	isMultiplier bool
}

type test struct {
	isDivisibleBy int
	targetTrue    int
	targetFalse   int
}

func (m *monkey) setStartingItems(input string) {
	var result []int64
	parts := strings.Split(input, ": ")
	items := strings.Split(parts[1], ", ")
	for _, parsedItem := range items {
		result = append(result, int64(utils.ToInt(parsedItem)))
	}
	m.items = result
}

func (m *monkey) setTest(detail string, input string) {
	parts := strings.Split(detail, "Test: divisible by ")
	t := test{}
	t.isDivisibleBy = utils.ToInt(parts[1])
	re := regexp.MustCompile(`[0-9]{1}`)
	throws := re.FindAllString(input, -1)
	t.targetTrue = utils.ToInt(throws[len(throws)-2])
	t.targetFalse = utils.ToInt(throws[len(throws)-1])
	m.test = t
}

func (m *monkey) setOperation(input string) {
	o := operation{}
	parts := strings.Split(input, ": new = old ")
	items := strings.Split(parts[1], " ")
	for i, item := range items {
		switch i {
		case 0:
			if item == "*" {
				o.isMultiplier = true
			} else {
				o.isMultiplier = false
			}
		case 1:
			if item == "old" {
				o.increasedBy = -1
			} else {
				o.increasedBy = int64(utils.ToInt(item))
			}
		default:
			continue
		}
	}
	m.operation = o
}

func (m *monkey) init(input string) {
	details := strings.Split(input, "\n")
	for i, detail := range details {
		switch i {
		case 1:
			m.setStartingItems(detail)
		case 2:
			m.setOperation(detail)
		case 3:
			m.setTest(detail, input)
		default:
			break
		}
	}
}

func (m *monkey) inspect(monkeys []monkey, relief bool) {
	d := 1
	for _, ms := range monkeys {
		d *= ms.test.isDivisibleBy
	}
	for _, item := range m.items {
		var worryLevel int64
		if m.operation.increasedBy == -1 {
			worryLevel = item * item
		}
		if m.operation.isMultiplier && m.operation.increasedBy != -1 {
			worryLevel = item * m.operation.increasedBy
		} else if !m.operation.isMultiplier && m.operation.increasedBy > 0 {
			worryLevel = item + m.operation.increasedBy
		}
		if relief {
			worryLevel = worryLevel / 3
		}
		var to *monkey
		worryLevel = worryLevel % int64(d)
		if worryLevel%int64(m.test.isDivisibleBy) == 0 {
			to = &monkeys[m.test.targetTrue]
		} else {
			to = &monkeys[m.test.targetFalse]
		}
		to.items = append(to.items, worryLevel)
	}
	m.inspectedItems += len(m.items)
	m.items = []int64{}

}

func parser(input string, f func([]monkey) string) string {
	var monkeys = []monkey{}
	parts := strings.Split(input, "\n\n")
	for _, part := range parts {
		m := monkey{}
		m.init(part)
		monkeys = append(monkeys, m)
	}

	return f(monkeys)
}

func getScore(monkeys []monkey) int {
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectedItems > monkeys[j].inspectedItems
	})
	return monkeys[0].inspectedItems * monkeys[1].inspectedItems
}

func solution1(monkeys []monkey) string {
	for i := 0; i < 20; i++ {
		for j := 0; j < len(monkeys); j++ {
			(&monkeys[j]).inspect(monkeys, true)
		}
	}
	return utils.ToStr(getScore(monkeys))
}

func solution2(monkeys []monkey) string {
	for i := 0; i < 10000; i++ {
		for j := 0; j < len(monkeys); j++ {
			(&monkeys[j]).inspect(monkeys, false)
		}
	}
	return utils.ToStr(getScore(monkeys))
}

func part1(input string) string {
	return parser(input, solution1)
}

func part2(input string) string {
	return parser(input, solution2)
}

// Run1 runs the part 1 solution for day 11
func (s Solution) Run1() common.Response {
	return common.Response{
		Output: part1(utils.ReadFile("./2022/day11/day11.txt")),
		Day:    "Day 11",
		Part:   "Part 1",
	}
}

// Run2 runs the part 2 solution for day 11
func (s Solution) Run2() common.Response {
	return common.Response{
		Output: part2(utils.ReadFile("./2022/day11/day11.txt")),
		Day:    "Day 11",
		Part:   "Part 2",
	}
}
