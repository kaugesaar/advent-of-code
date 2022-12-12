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
	"github.com/kaugesaar/advent-of-code/common"
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
func Run() {
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
	}

	fmt.Println("---- 🔨 Running benchmark 🔨 ----")
	fmt.Println("| Day | Part 1  | Part 2  |")
	fmt.Println("|-----|---------|---------|")
	day := 0

	for range b.queue {
		day++
		b.setFunction()
		part1 := testing.Benchmark(b.runBenchmark1)
		part2 := testing.Benchmark(b.runBenchmark2)
		if day < 10 {
			fmt.Printf("| %d   | %0.3fms ", day, toMs(part1.NsPerOp()))
		} else {
			fmt.Printf("| %d  | %0.3fms ", day, toMs(part1.NsPerOp()))
		}
		fmt.Printf("| %0.3fms |\n", toMs(part2.NsPerOp()))
	}
}

func (b *benchmark) setFunction() {
	b.f, b.queue = b.queue[0], append(b.queue[:0], b.queue[1:]...)
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
