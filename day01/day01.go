package day01

import (
	"strconv"
	"strings"
)

func ConvertInstructionsToNumbers(input string) []int {
	split := strings.Split(input, "\n")
	result := make([]int, len(split))
	for i, s := range split {
		result[i], _ = strconv.Atoi(s[1:])
		if s[0] == 'L' {
			result[i] = -result[i]
		}
	}
	return result
}

func CountPointToZero(input string) int {
	rotations := ConvertInstructionsToNumbers(input)
	current := 50
	count := 0
	for _, rotation := range rotations {
		current = (current + rotation) % 100
		if current == 0 {
			count++
		}
	}
	return count
}

func CountPointToZeroAfterAnyClick(input string) int {
	rotations := ConvertInstructionsToNumbers(input)
	current := 50
	count := 0
	for _, rotation := range rotations {
		count += AbsInt(rotation / 100)
		next := current + rotation%100
		if current != 0 && (next <= 0 || next >= 100) {
			count++
		}
		current = (next + 100) % 100
	}
	return count
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
