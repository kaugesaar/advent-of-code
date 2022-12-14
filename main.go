package main

import (
	"fmt"
	"os"
	"time"

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
	"github.com/kaugesaar/advent-of-code/2022/day16"
	"github.com/kaugesaar/advent-of-code/benchmark"
	"github.com/kaugesaar/advent-of-code/common"
	"github.com/kaugesaar/advent-of-code/utils"
)

// Solution can solve part1 and part2
type Solution interface {
	Run1() common.Response
	Run2() common.Response
}

var solutions = []Solution{
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
	day16.Solution{},
}

func run(run func() common.Response) {
	start := time.Now()
	response := run()
	elapsed := time.Since(start)
	ms := float64(elapsed.Nanoseconds()) / 1000000.0
	if response.Part == "Part 1" {
		fmt.Printf("⭐   : %v (%0.3fms)\n", response.Output, ms)
	} else {
		fmt.Printf("⭐⭐ : %v (%0.3fms)\n", response.Output, ms)
	}
}

func runSolution(day string, part string) {
	solution := solutions[utils.ToInt(day)-1]
	fmt.Printf("----- Day %v -----\n", day)
	switch part {
	case "1":
		run(solution.Run1)
	case "2":
		run(solution.Run2)
	default:
		run(solution.Run1)
		run(solution.Run2)
	}
}

func runAllSolutions() {
	for i, solution := range solutions {
		fmt.Printf("----- Day %v -----\n", i+1)
		run(solution.Run1)
		run(solution.Run2)
	}
}

func main() {
	numArgs := len(os.Args)

	switch numArgs {
	case 1:
		runAllSolutions()
	case 2:
		if os.Args[1] == "-bench" {
			benchmark.Run("")
		} else {
			runSolution(os.Args[1], "")
		}
	case 3:
		if os.Args[2] == "-bench" {
			benchmark.Run(os.Args[1])
		} else {
			runSolution(os.Args[1], os.Args[2])
		}
	default:
		panic("Invalid number of arguments")

	}
}
