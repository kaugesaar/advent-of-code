package day12

import (
	_ "embed"
	"testing"
)

//go:embed day12_test.txt
var testInput string

//go:embed day12.txt
var fileInput string

func TestPart1(t *testing.T) {
	result := part1(testInput)
	expected := "31"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := part2(testInput)
	expected := "29"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	result := part1(fileInput)
	expected := "481"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(fileInput)
	expected := "480"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(fileInput)
	}
}
func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2(fileInput)
	}
}
