package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input/day1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var sums []int
	var sum int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		kcal := strings.TrimSpace(scanner.Text())
		if kcal != "" {
			i, err := strconv.Atoi(kcal)
			if err != nil {
				panic(err)
			}
			sum += i
		} else {
			sums = append(sums, sum)
			sum = 0
		}
	}

	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})

	fmt.Println("Day 1 | Question 1:", sums[0])
	fmt.Println("Day 1 | Question 2:", sums[0]+sums[1]+sums[2])
}
