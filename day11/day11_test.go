package day11

import (
	"aoc2025/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetExample1() Graph {
	input := "aaa: you hhh\n" +
		"you: bbb ccc\n" +
		"bbb: ddd eee\n" +
		"ccc: ddd eee fff\n" +
		"ddd: ggg\n" +
		"eee: out\n" +
		"fff: out\n" +
		"ggg: out\n" +
		"hhh: ccc fff iii\n" +
		"iii: out"
	return ParseInput(input)
}

func TestExamplePart1(t *testing.T) {
	assert.Equal(t, 5, Part1(GetExample1()))
}

func TestPart1(t *testing.T) {
	input := helpers.Aoc().ReadPuzzleInputForTest(t, 11)
	assert.Equal(t, 574, Part1(ParseInput(input)))
}

func GetExample2() Graph {
	input := "svr: aaa bbb\n" +
		"aaa: fft\n" +
		"fft: ccc\n" +
		"bbb: tty\n" +
		"tty: ccc\n" +
		"ccc: ddd eee\n" +
		"ddd: hub\n" +
		"hub: fff\n" +
		"eee: dac\n" +
		"dac: fff\n" +
		"fff: ggg hhh\n" +
		"ggg: out\n" +
		"hhh: out"
	return ParseInput(input)
}

func TestExamplePart2(t *testing.T) {
	assert.Equal(t, 2, Part2(GetExample2()))
}

func TestPart2(t *testing.T) {
	input := helpers.Aoc().ReadPuzzleInputForTest(t, 11)
	assert.Equal(t, 306594217920240, Part2(ParseInput(input)))
}

func TestCreateDotFile(t *testing.T) {
	input := helpers.Aoc().ReadPuzzleInputForTest(t, 11)
	CreateDotFile(ParseInput(input))
}
