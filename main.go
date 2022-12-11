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

	if numArgs == 1 {
		runAllSolutions()
	}

	if numArgs == 2 {
		runSolution(os.Args[1], "")
	}

	if numArgs == 3 {
		runSolution(os.Args[1], os.Args[2])
	}
}
