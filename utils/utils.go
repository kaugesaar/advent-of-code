package utils

import (
	"io/ioutil"
	"math"
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

// MaxInt returns the smaller of a and b
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Abs returns the absolute value of x
func Abs(x int) int {
	return int(math.Abs(float64(x)))
}

// Contains checks if slice contains string
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
