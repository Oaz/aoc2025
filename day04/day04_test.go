package day04

import (
	"aoc2025/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetExample() Map {
	input := "..@@.@@@@.\n" +
		"@@@.@.@.@@\n" +
		"@@@@@.@.@@\n" +
		"@.@@@@..@.\n" +
		"@@.@@@@.@@\n" +
		".@@@@@@@.@\n" +
		".@.@.@.@@@\n" +
		"@.@@@.@@@@\n" +
		".@@@@@@@@.\n" +
		"@.@.@@@.@."
	return InputToMap(input)
}

func TestExamplePart1(t *testing.T) {
	forklifts := GetExample()
	assert.Equal(t, 3, CountNeighbors(forklifts, Coordinates{X: 3, Y: 0}))
	assert.Equal(t, 8, CountNeighbors(forklifts, Coordinates{X: 4, Y: 4}))
	assert.Equal(t, 13, CountReachableForklifts(forklifts))
}

func TestPart1(t *testing.T) {
	input := helpers.Aoc().ReadPuzzleInputForTest(t, 4)
	assert.Equal(t, 1493, CountReachableForklifts(InputToMap(input)))
}

func TestExamplePart2(t *testing.T) {
	assert.Equal(t, 43, CountAllReachableForklifts(GetExample()))
}

func TestPart2(t *testing.T) {
	input := helpers.Aoc().ReadPuzzleInputForTest(t, 4)
	assert.Equal(t, 9194, CountAllReachableForklifts(InputToMap(input)))
}
