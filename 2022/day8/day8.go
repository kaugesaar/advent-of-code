package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/kaugesaar/advent-of-code/utils"
)

//go:embed day8.txt
var fileInput string

func parser(input string, s func([][]int) int) int {
	var trees [][]int
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		var r []int
		for _, char := range row {
			r = append(r, utils.ToInt(string(char)))
		}
		trees = append(trees, r)
	}
	return s(trees)
}

func isVisibleRow(trees []int, col int, height int) bool {
	for _, row := range trees {
		if row >= height {
			return false
		}
	}
	return true
}

func isVisibleCol(trees [][]int, col int, height int) bool {
	for _, row := range trees {
		if row[col] >= height {
			return false
		}
	}
	return true
}

func transpose(trees [][]int, col int) []int {
	var res []int
	for row := range trees {
		res = append(res, trees[row][col])
	}
	return res
}

func parseScore(ctx *[]int, trees []int) {
	score := 0
	for _, i := range trees {
		score++
		if i >= (*ctx)[1] {
			break
		}
	}
	(*ctx)[0] *= score
}

func parseScoreReverse(ctx *[]int, trees []int) {
	score := 0
	for i := len(trees) - 1; i >= 0; i-- {
		score++
		if trees[i] >= (*ctx)[1] {
			break
		}
	}
	(*ctx)[0] *= score
}

func solution1(trees [][]int) int {
	var count int
	for row := range trees {
		for col, height := range trees[row] {
			if isVisibleCol(trees[:row], col, height) ||
				isVisibleRow(trees[row][:col], col, height) ||
				isVisibleCol(trees[row+1:], col, height) ||
				isVisibleRow(trees[row][col+1:], col, height) {
				count++
			}
		}
	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solution2(trees [][]int) int {
	var highScore int
	for row := 0; row < len(trees); row++ {
		for col := 0; col < len(trees[0]); col++ {
			ctx := []int{1, trees[row][col]}
			parseScoreReverse(&ctx, (transpose(trees[:row], col)))
			parseScoreReverse(&ctx, trees[row][:col])
			parseScore(&ctx, transpose(trees[row+1:], col))
			parseScore(&ctx, trees[row][col+1:])
			highScore = max(highScore, ctx[0])

		}
	}
	return highScore
}

func part1(input string) int {
	return parser(input, solution1)
}

func part2(input string) int {
	return parser(input, solution2)
}

func main() {
	fmt.Println("---- 2022 Day 8 ----")
	fmt.Println("part1: ", part1(fileInput))
	fmt.Println("part2: ", part2(fileInput))
}
