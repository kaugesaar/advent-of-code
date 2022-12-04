package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input/day4.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		input = append(input, text)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Day 2 | Question 1:", questionOne(input))
	fmt.Println("Day 2 | Question 1:", questionTwo(input))
}

func questionOne(input []string) int {
	var count int
	for _, pairs := range input {
		a, b := parsePair(pairs)
		if fullyContains(a, b) {
			count++
		}
	}
	return count
}

func questionTwo(input []string) int {
	var count int
	for _, pairs := range input {
		a, b := parsePair(pairs)
		if overlaps(a, b) {
			count++
		}
	}
	return count
}

func parsePair(pairs string) ([]int, []int) {
	assignments := strings.Split(pairs, ",")
	a := parseSection(assignments[0])
	b := parseSection(assignments[1])
	return a, b
}

func parseSection(assignments string) []int {
	sections := strings.Split(assignments, "-")
	low, _ := strconv.Atoi(sections[0])
	high, _ := strconv.Atoi(sections[1])
	return []int{low, high}
}

func fullyContains(a []int, b []int) bool {
	aLow := a[0]
	aHigh := a[1]
	bLow := b[0]
	bHigh := b[1]
	return (aLow <= bLow && aHigh >= bHigh) || (bLow <= aLow && bHigh >= aHigh)
}

func overlaps(a []int, b []int) bool {
	aLow := a[0]
	aHigh := a[1]
	bLow := b[0]
	bHigh := b[1]
	return (aLow <= bLow && bLow <= aHigh) || (bLow <= aLow && aLow <= bHigh)
}
