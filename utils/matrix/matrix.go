package matrix

import (
	"fmt"
	"strings"
)

// ByteMatrix d
type ByteMatrix [][]byte

// IntMatrix d
type IntMatrix [][]int

// Pos position
type Pos struct{ X, Y int }

// String formats Pos as a string
func (p Pos) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

// BuildByteMatrix returns a matrix of bytes. Takes a string array as input.
func BuildByteMatrix(lines []string) ByteMatrix {
	m := make([][]uint8, len(lines))
	for j, l := range lines {
		l = strings.TrimSpace(l)
		m[j] = make([]byte, len(l))
		for i, c := range l {
			m[j][i] = byte(c - '0')
		}
	}
	return m
}

// BuildIntMatrix returns a matrix of ints. Takes a string array as input.
func BuildIntMatrix(lines []string) IntMatrix {
	m := make([][]int, len(lines))
	for j, l := range lines {
		l = strings.TrimSpace(l)
		m[j] = make([]int, len(l))
		for i, c := range l {
			m[j][i] = int(c)
		}
	}
	return m
}
