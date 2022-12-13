package day13

import (
	_ "embed"
	"testing"
)

//go:embed day13_test.txt
var testInput string

//go:embed day13.txt
var fileInput string

func TestPart1(t *testing.T) {
	result := part1(testInput)
	expected := "13"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := part2(testInput)
	expected := "140"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	result := part1(fileInput)
	expected := "5529"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(fileInput)
	expected := "27690"
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
