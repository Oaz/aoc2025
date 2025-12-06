package day06

import (
	"aoc2025/helpers"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetExample() string {
	return "123 328  51 64 \n" +
		" 45 64  387 23 \n" +
		"  6 98  215 314\n" +
		"*   +   *   +  "
}

func TestExamplePart1(t *testing.T) {
	homework := InputToMathHomework(GetExample())
	assert.Equal(t, 33210, homework[0].Solve())
	assert.Equal(t, 490, homework[1].Solve())
	assert.Equal(t, 4243455, homework[2].Solve())
	assert.Equal(t, 401, homework[3].Solve())
	assert.Equal(t, 4277556, homework.GrandTotal())
}

func TestPart1(t *testing.T) {
	input := helpers.Aoc().ReadPuzzleInputForTest(t, 6)
	assert.Equal(t, 6295830249262, InputToMathHomework(input).GrandTotal())
}

func TestTranspose(t *testing.T) {
	matrix := strings.Split(GetExample(), "\n")
	matrix = matrix[0 : len(matrix)-1]
	assert.Equal(t, []string{
		"1  ", "24 ", "356", "   ",
		"369", "248", "8  ", "   ",
		" 32", "581", "175", "   ",
		"623", "431", "  4",
	}, Transpose(matrix))
}

func TestTransposeAndConvert(t *testing.T) {
	matrix := strings.Split(GetExample(), "\n")
	matrix = matrix[0 : len(matrix)-1]
	assert.Equal(t, [][]int{
		{1, 24, 356},
		{369, 248, 8},
		{32, 581, 175},
		{623, 431, 4},
	}, TransposeAndConvert(matrix))
}

func TestExamplePart2(t *testing.T) {
	homework := InputTransposedToMathHomework(GetExample())
	assert.Equal(t, 8544, homework[0].Solve())
	assert.Equal(t, 625, homework[1].Solve())
	assert.Equal(t, 3253600, homework[2].Solve())
	assert.Equal(t, 1058, homework[3].Solve())
	assert.Equal(t, 3263827, homework.GrandTotal())
}

func TestPart2(t *testing.T) {
	input := helpers.Aoc().ReadPuzzleInputForTest(t, 6)
	assert.Equal(t, 9194682052782, InputTransposedToMathHomework(input).GrandTotal())
}
