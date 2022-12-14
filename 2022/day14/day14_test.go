package day14

import (
	_ "embed"
	"testing"
)

//go:embed day14_test.txt
var testInput string

//go:embed day14.txt
var fileInput string

func TestPart1(t *testing.T) {
	result := part1(testInput)
	expected := "24"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := part2(testInput)
	expected := "93"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	result := part1(fileInput)
	expected := "774"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(fileInput)
	expected := "22195"
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
