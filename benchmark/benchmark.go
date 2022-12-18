package benchmark

import (
	"fmt"
	"testing"

	"github.com/kaugesaar/advent-of-code/2022/day01"
	"github.com/kaugesaar/advent-of-code/2022/day02"
	"github.com/kaugesaar/advent-of-code/2022/day03"
	"github.com/kaugesaar/advent-of-code/2022/day04"
	"github.com/kaugesaar/advent-of-code/2022/day05"
	"github.com/kaugesaar/advent-of-code/2022/day06"
	"github.com/kaugesaar/advent-of-code/2022/day07"
	"github.com/kaugesaar/advent-of-code/2022/day08"
	"github.com/kaugesaar/advent-of-code/2022/day09"
	"github.com/kaugesaar/advent-of-code/2022/day10"
	"github.com/kaugesaar/advent-of-code/2022/day11"
	"github.com/kaugesaar/advent-of-code/2022/day12"
	"github.com/kaugesaar/advent-of-code/2022/day13"
	"github.com/kaugesaar/advent-of-code/2022/day14"
	"github.com/kaugesaar/advent-of-code/2022/day15"
	"github.com/kaugesaar/advent-of-code/common"
	"github.com/kaugesaar/advent-of-code/utils"
)

// Solution can solve part1 and part2
type Solution interface {
	Run1() common.Response
	Run2() common.Response
}

type benchmark struct {
	f     Solution
	queue []Solution
}

// Run all benchmarks
func Run(solution string) {
	var b = benchmark{}
	b.queue = []Solution{
		day01.Solution{},
		day02.Solution{},
		day03.Solution{},
		day04.Solution{},
		day05.Solution{},
		day06.Solution{},
		day07.Solution{},
		day08.Solution{},
		day09.Solution{},
		day10.Solution{},
		day11.Solution{},
		day12.Solution{},
		day13.Solution{},
		day14.Solution{},
		day15.Solution{},
	}

	day := 1

	if solution != "" {
		day = utils.ToInt(solution)
		b.queue = []Solution{b.queue[utils.ToInt(solution) - 1]}
	}

	fmt.Println("---- 🔨 Running benchmark 🔨 ----")
	fmt.Println("| Day | Part 1  | Part 2  |")
	fmt.Println("|-----|---------|---------|")

	for _, f := range b.queue {
		b.setFunction(f)
		part1 := testing.Benchmark(b.runBenchmark1)
		part2 := testing.Benchmark(b.runBenchmark2)
		if day < 10 {
			fmt.Printf("| %d   | %.3fms ", day, toMs(part1.NsPerOp()))
		} else {
			fmt.Printf("| %d  | %.3fms ", day, toMs(part1.NsPerOp()))
		}
		fmt.Printf("| %.3fms |\n", toMs(part2.NsPerOp()))
		day++
	}
}

func (b *benchmark) setFunction(f Solution) {
	b.f = f
}

func (b *benchmark) runBenchmark1(t *testing.B) {
	for i := 0; i < t.N; i++ {
		b.f.Run1()
	}
}

func (b *benchmark) runBenchmark2(t *testing.B) {
	for i := 0; i < t.N; i++ {
		b.f.Run2()
	}
}

func toMs(nsPerOp int64) float64 {
	return float64(nsPerOp) / 1000000.0
}
