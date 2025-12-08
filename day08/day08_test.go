package day08

import (
	"aoc2025/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetExample() []Box {
	input := "162,817,812\n" +
		"57,618,57\n" +
		"906,360,560\n" +
		"592,479,940\n" +
		"352,342,300\n" +
		"466,668,158\n" +
		"542,29,236\n" +
		"431,825,988\n" +
		"739,650,466\n" +
		"52,470,668\n" +
		"216,146,977\n" +
		"819,987,18\n" +
		"117,168,530\n" +
		"805,96,715\n" +
		"346,949,466\n" +
		"970,615,88\n" +
		"941,993,340\n" +
		"862,61,35\n" +
		"984,92,344\n" +
		"425,690,689"
	return ParseInput(input)
}

func TestExamplePart1(t *testing.T) {
	assert.Equal(t, 40, Part1(GetExample(), 10))
}

func TestPart1(t *testing.T) {
	input := helpers.Aoc().ReadPuzzleInputForTest(t, 8)
	assert.Equal(t, 175500, Part1(ParseInput(input), 1000))
}

func TestExamplePart2(t *testing.T) {
	assert.Equal(t, 25272, Part2(GetExample()))
}

func TestPart2(t *testing.T) {
	input := helpers.Aoc().ReadPuzzleInputForTest(t, 8)
	assert.Equal(t, 6934702555, Part2(ParseInput(input)))
}
