package main

import (
	_ "embed"
	"testing"
)

//go:embed day7_test.txt
var testInput string

func TestPart1(t *testing.T) {
	result := parser(testInput, solution1)
	expected := 95437
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := parser(testInput, solution2)
	expected := 24933642
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	result := parser(fileInput, solution1)
	expected := 1391690
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := parser(fileInput, solution2)
	expected := 5469168
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parser(testInput, solution1)
	}
}
func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parser(testInput, solution2)
	}
}
