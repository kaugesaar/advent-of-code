package utils

import "strconv"

// ToInt transforms a string to an integer
func ToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}
