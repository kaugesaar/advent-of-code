package main

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
)

var benches = []func(b *testing.B){
	benchmarkDay1Part1,
	benchmarkDay1Part2,
	benchmarkDay2Part1,
	benchmarkDay2Part2,
	benchmarkDay3Part1,
	benchmarkDay3Part2,
	benchmarkDay4Part1,
	benchmarkDay4Part2,
	benchmarkDay5Part1,
	benchmarkDay5Part2,
	benchmarkDay6Part1,
	benchmarkDay6Part2,
	benchmarkDay7Part1,
	benchmarkDay7Part2,
	benchmarkDay8Part1,
	benchmarkDay8Part2,
}

func main() {
	fmt.Println("---- 🔨 Running benchmark 🔨 ----")
	fmt.Println("| Day | Part 1  | Part 2  |")
	fmt.Println("|-----|---------|---------|")
	day := 1
	for i, bench := range benches {
		testing := testing.Benchmark(bench)
		if i%2 != 0 {
			fmt.Printf("| %0.3fms |\n", toMs(testing.NsPerOp()))
			day++
		} else {
			if day < 10 {
				fmt.Printf("| %d   | %0.3fms ", day, toMs(testing.NsPerOp()))
			} else {
				fmt.Printf("| %d  | %0.3fms ", day, toMs(testing.NsPerOp()))
			}
		}
	}
}

func toMs(nsPerOp int64) float64 {
	return float64(nsPerOp) / 1000000.0
}

func benchmarkDay1Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day01.Solution{}.Run1()
	}
}

func benchmarkDay1Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day01.Solution{}.Run2()
	}
}

func benchmarkDay2Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day02.Solution{}.Run1()
	}
}

func benchmarkDay2Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day02.Solution{}.Run2()
	}
}

func benchmarkDay3Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day03.Solution{}.Run1()
	}
}

func benchmarkDay3Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day03.Solution{}.Run2()
	}
}

func benchmarkDay4Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day04.Solution{}.Run1()
	}
}

func benchmarkDay4Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day04.Solution{}.Run2()
	}
}

func benchmarkDay5Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day05.Solution{}.Run1()
	}
}

func benchmarkDay5Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day05.Solution{}.Run2()
	}
}

func benchmarkDay6Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day06.Solution{}.Run1()
	}
}

func benchmarkDay6Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day06.Solution{}.Run2()
	}
}

func benchmarkDay7Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day07.Solution{}.Run1()
	}
}

func benchmarkDay7Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day07.Solution{}.Run2()
	}
}

func benchmarkDay8Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day08.Solution{}.Run1()
	}
}

func benchmarkDay8Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day08.Solution{}.Run2()
	}
}
