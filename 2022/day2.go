package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func questionOne(games []string) int {
	var sum int
	for _, game := range games {
		switch strings.TrimSpace(game) {
		case "A X":
			sum += 4
		case "A Y":
			sum += 8
		case "A Z":
			sum += 3
		case "B X":
			sum++
		case "B Y":
			sum += 5
		case "B Z":
			sum += 9
		case "C X":
			sum += 7
		case "C Y":
			sum += 2
		case "C Z":
			sum += 6
		default:
		}
	}
	return sum
}

func questionTwo(games []string) int {
	var sum int
	for _, game := range games {
		switch strings.TrimSpace(game) {
		case "A X":
			sum += 3
		case "A Y":
			sum += 4
		case "A Z":
			sum += 8
		case "B X":
			sum++
		case "B Y":
			sum += 5
		case "B Z":
			sum += 9
		case "C X":
			sum += 2
		case "C Y":
			sum += 6
		case "C Z":
			sum += 7
		default:
		}
	}
	return sum
}

func main() {
	file, err := os.Open("./input/day2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []string

	for scanner.Scan() {
		text := scanner.Text()
		input = append(input, text)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Day 2 | Question 1:", questionOne(input))
	fmt.Println("Day 2 | Question 2:", questionTwo(input))
}
