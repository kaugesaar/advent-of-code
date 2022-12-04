package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input/day3.txt")
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

func questionOne(rucksacks []string) int {
	var sum int
	for _, rucksack := range rucksacks {
		a := rucksack[:len(rucksack)/2]
		b := rucksack[len(rucksack)/2:]
		for _, c := range a {
			if strings.ContainsRune(b, c) {
				sum += score(c)
				break
			}
		}
	}
	return sum
}

func questionTwo(rucksacks []string) int {
	var sum int
	for i := 0; i < len(rucksacks); i += 3 {
		r1 := rucksacks[i]
		r2 := rucksacks[i+1]
		r3 := rucksacks[i+2]
		for _, c := range r1 {
			if strings.ContainsRune(r2, c) && strings.ContainsRune(r3, c) {
				sum += score(c)
				break
			}
		}
	}
	return sum
}

func score(item rune) int {
	switch {
	case 'A' <= item && item <= 'Z':
		return int(item-'A') + 27
	case 'a' <= item && item <= 'z':
		return int(item-'a') + 1
	default:
		return 0
	}
}
