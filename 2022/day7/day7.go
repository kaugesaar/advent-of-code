package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/kaugesaar/advent-of-code/utils"
)

//go:embed day7.txt
var fileInput string

func parser(input string, s func([]int) int) int {
	lines := strings.Split(input, "\n")

	var dir []int
	var directories []int

	for _, line := range lines {
		if strings.HasPrefix(line, "$ cd") {
			if strings.Contains(line, "..") {
				directories = append(directories, dir[len(dir)-1])
				dir = dir[:len(dir)-1]
			} else {
				dir = append(dir, 0)
			}
		}
		if !strings.HasPrefix(line, "$") && !strings.HasPrefix(line, "dir") {
			fields := strings.Fields(line)
			size := utils.ToInt(fields[0])
			for i := range dir {
				dir[i] += size
			}
		}
	}
	directories = append(directories, dir...)
	return s(directories)
}

func solution1(directories []int) int {
	var sum int
	for _, size := range directories {
		if size <= 100000 {
			sum += size
		}
	}
	return sum
}

func solution2(directories []int) int {
	max := 0
	for _, size := range directories {
		if size > max {
			max = size
		}
	}
	req := 30000000 - (70000000 - max)
	min := 0
	for _, size := range directories {
		if size >= req && (min == 0 || size < min) {
			min = size
		}
	}
	return min
}

func main() {
	fmt.Println("---- 2022 Day 7 ----")
	fmt.Println("part1: ", parser(fileInput, solution1))
	fmt.Println("part2: ", parser(fileInput, solution2))
}
