package utils

import (
	"io/ioutil"
	"strconv"
)

// ReadFile reads the named file and returns the content.
func ReadFile(path string) string {
	input, error := ioutil.ReadFile(path)
	if error != nil {
		panic(error)
	}
	return string(input)
}

// ToInt transforms a string to an integer
func ToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

// ToStr transforms a integer to a string
func ToStr(i int) string {
	str := strconv.Itoa(i)
	return str
}

// MinInt returns the smaller of a and b
func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// MinInt returns the smaller of a and b
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

