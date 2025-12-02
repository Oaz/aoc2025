package day02

import (
	"aoc2025/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetExample() []Range {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	return InputToListOfRanges(input)
}

func TestExamplePart1(t *testing.T) {
	ranges := GetExample()
	policy := RepeatTwicePolicy()
	assert.Equal(t, []int{11, 22}, FindFakeIDs(ranges[0], policy))
	assert.Equal(t, []int{99}, FindFakeIDs(ranges[1], policy))
	assert.Equal(t, []int{1010}, FindFakeIDs(ranges[2], policy))
	assert.Equal(t, []int{1188511885}, FindFakeIDs(ranges[3], policy))
	assert.Equal(t, []int{222222}, FindFakeIDs(ranges[4], policy))
	assert.Equal(t, []int{}, FindFakeIDs(ranges[5], policy))
	assert.Equal(t, []int{446446}, FindFakeIDs(ranges[6], policy))
	assert.Equal(t, []int{38593859}, FindFakeIDs(ranges[7], policy))
	assert.Equal(t, []int{}, FindFakeIDs(ranges[8], policy))
	assert.Equal(t, []int{}, FindFakeIDs(ranges[9], policy))
	assert.Equal(t, []int{}, FindFakeIDs(ranges[10], policy))
	assert.Equal(t, 1227775554, SumFakeIDs(ranges, policy))
}

func TestPart1(t *testing.T) {
	input := helpers.Aoc().ReadPuzzleInputForTest(t, 2)
	assert.Equal(t, 29940924880, SumFakeIDs(InputToListOfRanges(input), RepeatTwicePolicy()))
}

func TestExamplePart2(t *testing.T) {
	ranges := GetExample()
	policy := RepeatAtLeastTwicePolicy()
	assert.Equal(t, []int{11, 22}, FindFakeIDs(ranges[0], policy))
	assert.Equal(t, []int{99, 111}, FindFakeIDs(ranges[1], policy))
	assert.Equal(t, []int{999, 1010}, FindFakeIDs(ranges[2], policy))
	assert.Equal(t, []int{1188511885}, FindFakeIDs(ranges[3], policy))
	assert.Equal(t, []int{222222}, FindFakeIDs(ranges[4], policy))
	assert.Equal(t, []int{}, FindFakeIDs(ranges[5], policy))
	assert.Equal(t, []int{446446}, FindFakeIDs(ranges[6], policy))
	assert.Equal(t, []int{38593859}, FindFakeIDs(ranges[7], policy))
	assert.Equal(t, []int{565656}, FindFakeIDs(ranges[8], policy))
	assert.Equal(t, []int{824824824}, FindFakeIDs(ranges[9], policy))
	assert.Equal(t, []int{2121212121}, FindFakeIDs(ranges[10], policy))
	assert.Equal(t, 4174379265, SumFakeIDs(ranges, policy))
}

func TestPart2(t *testing.T) {
	input := helpers.Aoc().ReadPuzzleInputForTest(t, 2)
	assert.Equal(t, 48631958998, SumFakeIDs(InputToListOfRanges(input), RepeatAtLeastTwicePolicy()))
}
