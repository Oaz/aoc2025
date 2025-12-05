package day05

import (
	"aoc2025/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetExample() Inventory {
	input :=
		"3-5\n" +
			"10-14\n" +
			"16-20\n" +
			"12-18\n" +
			"\n" +
			"1\n" +
			"5\n" +
			"8\n" +
			"11\n" +
			"17\n" +
			"32"
	return InputToInventory(input)
}

func TestExamplePart1(t *testing.T) {
	assert.Equal(t, 3, GetExample().CountFreshAvailableIngredients())
}

func TestPart1(t *testing.T) {
	input := helpers.Aoc().ReadPuzzleInputForTest(t, 5)
	assert.Equal(t, 821, InputToInventory(input).CountFreshAvailableIngredients())
}

func TestExamplePart2(t *testing.T) {
	assert.Equal(t, 14, GetExample().CountFreshIngredients())
}

func TestPart2(t *testing.T) {
	input := helpers.Aoc().ReadPuzzleInputForTest(t, 5)
	assert.Equal(t, 344771884978261, InputToInventory(input).CountFreshIngredients())
}
