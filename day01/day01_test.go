package day01

import (
	"aoc2025/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExamplePart1(t *testing.T) {
	input := "L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82"
	assert.Equal(t, []int{-68, -30, 48, -5, 60, -55, -1, -99, 14, -82}, ConvertInstructionsToNumbers(input))
	assert.Equal(t, 3, CountPointToZero(input))
}

func TestPart1(t *testing.T) {
	input := helpers.Aoc().ReadPuzzleInputForTest(t, 1)
	password := CountPointToZero(input)
	assert.Equal(t, 1141, password)
}

func TestExamplePart2(t *testing.T) {
	input := "L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82"
	assert.Equal(t, 6, CountPointToZeroAfterAnyClick(input))
}

func TestCarefulExamplePart2(t *testing.T) {
	input := "L168\nL230\nR48\nL5\nR560\nL55\nL1001\nL99\nR14\nL82"
	assert.Equal(t, 24, CountPointToZeroAfterAnyClick(input))
}

func TestPart2(t *testing.T) {
	input := helpers.Aoc().ReadPuzzleInputForTest(t, 1)
	password := CountPointToZeroAfterAnyClick(input)
	assert.Equal(t, 6634, password)
}
