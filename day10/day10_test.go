package day10

import (
	"aoc2025/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetExample() []Machine {
	input := "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}\n" +
		"[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}\n" +
		"[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}"
	return ParseInput(input)
}

func TestExampleParsing(t *testing.T) {
	machines := GetExample()
	assert.Equal(t, Machine{
		LightsTarget: 6,
		LightButtons: []int{8, 10, 4, 12, 5, 3},
		Buttons:      [][]int{{3}, {1, 3}, {2}, {2, 3}, {0, 2}, {0, 1}},
		Joltages:     []uint{3, 5, 4, 7},
	}, machines[0])
	assert.Equal(t, Machine{
		LightsTarget: 8,
		LightButtons: []int{29, 12, 17, 7, 30},
		Buttons:      [][]int{{0, 2, 3, 4}, {2, 3}, {0, 4}, {0, 1, 2}, {1, 2, 3, 4}},
		Joltages:     []uint{7, 5, 12, 7, 2},
	}, machines[1])
	assert.Equal(t, Machine{
		LightsTarget: 46,
		LightButtons: []int{31, 25, 55, 6},
		Buttons:      [][]int{{0, 1, 2, 3, 4}, {0, 3, 4}, {0, 1, 2, 4, 5}, {1, 2}},
		Joltages:     []uint{10, 11, 11, 5, 10, 5},
	}, machines[2])
}

func TestExampleButtonPressesForLights(t *testing.T) {
	machines := GetExample()
	assert.Equal(t, machines[0].LightsTarget, machines[0].LightsAfterButtonPresses([]int{0, 1, 2}))
	assert.Equal(t, machines[0].LightsTarget, machines[0].LightsAfterButtonPresses([]int{1, 3, 5, 5}))
	assert.Equal(t, machines[0].LightsTarget, machines[0].LightsAfterButtonPresses([]int{0, 2, 3, 4, 5}))
	assert.Equal(t, machines[0].LightsTarget, machines[0].LightsAfterButtonPresses([]int{4, 5}))
	assert.Equal(t, machines[1].LightsTarget, machines[1].LightsAfterButtonPresses([]int{2, 3, 4}))
	assert.Equal(t, machines[2].LightsTarget, machines[2].LightsAfterButtonPresses([]int{1, 2}))
}

func TestExamplePart1(t *testing.T) {
	assert.Equal(t, 7, Part1(GetExample()))
}

func TestPart1(t *testing.T) {
	input := helpers.Aoc().ReadPuzzleInputForTest(t, 10)
	assert.Equal(t, 571, Part1(ParseInput(input)))
}

func TestExamplePart2(t *testing.T) {
	assert.Equal(t, 33, Part2(GetExample()))
}

func TestBigExample1(t *testing.T) {
	input := "[#.##.#...] (0,1,2,3,5,7,8) (0,1,2,3,4,6,7,8) (1) (0,2,3,4,5,6,7) (0,1,7) (0,2,3,5) (3,8) (1,6) (0,1,2,4) (0,1,6,7,8) {61,67,34,31,28,23,53,50,24}"
	assert.Equal(t, 92, Part2(ParseInput(input)))
}

func TestBigExample2(t *testing.T) {
	input := "[#####....#] (1,4,8) (0,1,2,7,8) (4,6,7) (0,4,5,6,8) (1,3,6,7,9) (0,5,6,8) (4,6) (1,3,8,9) (0,4,9) (0,1,3,4,5,7,8,9) (0,2,3,7,9) (1,2,3) {39,55,32,44,42,17,48,43,48,34}"
	assert.Equal(t, 99, Part2(ParseInput(input)))
}

func TestBigExample3(t *testing.T) {
	input := "[#.#..#.#.] (0,3,4,5,7,8) (3,5,8) (1,2,3,7,8) (0,1,6) (0,1,2,3,7) (0,3,4,5,6,7,8) (1,2,3,4,5,6,7,8) (0,2,3,4,5,6,8) (1,3,5,6,7) (1,3,4,5,7) (1,4,6,7) {47,79,34,83,57,68,74,77,68}"
	assert.Equal(t, 114, Part2(ParseInput(input)))
}

func TestBigExample4(t *testing.T) {
	input := "[##....##..] (0,2,3,4,6,7,8) (2,7,8) (0,1,2,3,4,6,7,8) (0,2,3,4,5,7,9) (0,1,6,7) (0,1,2,3,5,6,8,9) (0,1,3,5,7,8,9) (0,2,7,8,9) (0,2,3,7,8,9) (6,8) (0,1,2,3,5,6,9) (0,1,2,3,5,6,7,8,9) {214,185,195,192,21,180,184,193,182,196}"
	assert.Equal(t, 222, Part2(ParseInput(input)))
}
func TestPart2(t *testing.T) {
	input := helpers.Aoc().ReadPuzzleInputForTest(t, 10)
	assert.Equal(t, 20869, Part2(ParseInput(input)))
}
