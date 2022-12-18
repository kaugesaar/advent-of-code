package day15

import (
	_ "embed"
	"testing"
)

//go:embed day15_test.txt
var testInput string

//go:embed day15.txt
var fileInput string

func TestPart1(t *testing.T) {
	result := part1(testInput, 10)
	expected := "26"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := part2(testInput, 20)
	expected := "56000011"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	result := part1(fileInput, 2000000)
	expected := "4748135"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(fileInput, 4000000)
	expected := "13743542639657"
	if result != expected {
		t.Errorf("Result is incorrect, got: %s, want: %s.", result, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(fileInput, 2000000)
	}
}
func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2(fileInput, 4000000)
	}
}
