package day03

import (
	"aoc2025/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetExample() []Bank {
	input := "987654321111111\n811111111111119\n234234234234278\n818181911112111\n"
	return InputToBanks(input)
}

func TestExamplePart1(t *testing.T) {
	banks := GetExample()
	assert.Equal(t, 98, banks[0].Joltage(2))
	assert.Equal(t, 89, banks[1].Joltage(2))
	assert.Equal(t, 78, banks[2].Joltage(2))
	assert.Equal(t, 92, banks[3].Joltage(2))
	assert.Equal(t, 357, TotalOutputJoltage(banks, 2))
}

func TestPart1(t *testing.T) {
	input := helpers.Aoc().ReadPuzzleInputForTest(t, 3)
	assert.Equal(t, 17193, TotalOutputJoltage(InputToBanks(input), 2))
}

func TestExamplePart2(t *testing.T) {
	banks := GetExample()
	assert.Equal(t, 987654321111, banks[0].Joltage(12))
	assert.Equal(t, 811111111119, banks[1].Joltage(12))
	assert.Equal(t, 434234234278, banks[2].Joltage(12))
	assert.Equal(t, 888911112111, banks[3].Joltage(12))
	assert.Equal(t, 3121910778619, TotalOutputJoltage(banks, 12))
}

func TestPart2(t *testing.T) {
	input := helpers.Aoc().ReadPuzzleInputForTest(t, 3)
	assert.Equal(t, 171297349921310, TotalOutputJoltage(InputToBanks(input), 12))
}
