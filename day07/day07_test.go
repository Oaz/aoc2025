package day07

import (
	"aoc2025/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetExample() Manifold {
	input :=
		".......S.......\n" +
			"...............\n" +
			".......^.......\n" +
			"...............\n" +
			"......^.^......\n" +
			"...............\n" +
			".....^.^.^.....\n" +
			"...............\n" +
			"....^.^...^....\n" +
			"...............\n" +
			"...^.^...^.^...\n" +
			"...............\n" +
			"..^...^.....^..\n" +
			"...............\n" +
			".^.^.^.^.^...^.\n" +
			"..............."
	return InputToManifold(input)
}

func TestExamplePart1(t *testing.T) {
	assert.Equal(t, 21, GetExample().CountSplits())
}

func TestPart1(t *testing.T) {
	input := helpers.Aoc().ReadPuzzleInputForTest(t, 7)
	assert.Equal(t, 1698, InputToManifold(input).CountSplits())
}

func TestExamplePart2(t *testing.T) {
	assert.Equal(t, 40, GetExample().CountTimelines())
}

func TestPart2(t *testing.T) {
	input := helpers.Aoc().ReadPuzzleInputForTest(t, 7)
	assert.Equal(t, 95408386769474, InputToManifold(input).CountTimelines())
}
