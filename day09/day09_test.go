package day09

import (
	"aoc2025/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetExample() Floor {
	input := "7,1\n" +
		"11,1\n" +
		"11,7\n" +
		"9,7\n" +
		"9,5\n" +
		"2,5\n" +
		"2,3\n" +
		"7,3"
	return ParseInput(input)
}

func TestExamplePart1(t *testing.T) {
	assert.Equal(t, 50, Part1(GetExample()))
}

func TestPart1(t *testing.T) {
	input := helpers.Aoc().ReadPuzzleInputForTest(t, 9)
	assert.Equal(t, 4729332959, Part1(ParseInput(input)))
}

func TestExamplePart2(t *testing.T) {
	assert.Equal(t, 24, Part2(GetExample()))
}

func TestPart2(t *testing.T) {
	input := helpers.Aoc().ReadPuzzleInputForTest(t, 9)
	assert.Equal(t, 1474477524, Part2(ParseInput(input)))
}
