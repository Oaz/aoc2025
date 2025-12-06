package day06

import (
	"strconv"
	"strings"
)

type Operation interface {
	Identity() int
	Reduce(int, int) int
	String() string
}

type NoOp struct{}

func (NoOp) Identity() int       { return 0 }
func (NoOp) Reduce(_, _ int) int { return 0 }
func (NoOp) String() string      { return "" }

type AddOp struct{}

func (AddOp) Identity() int       { return 0 }
func (AddOp) Reduce(a, b int) int { return a + b }
func (AddOp) String() string      { return "+" }

type MulOp struct{}

func (MulOp) Identity() int       { return 1 }
func (MulOp) Reduce(a, b int) int { return a * b }
func (MulOp) String() string      { return "*" }

func NewOperation(op string) Operation {
	switch op {
	case "+":
		return AddOp{}
	case "*":
		return MulOp{}
	}
	return NoOp{}
}

type Problem struct {
	Operands []int
	Operator Operation
}

func (problem Problem) Solve() int {
	operator := problem.Operator
	result := operator.Identity()
	for _, operand := range problem.Operands {
		result = operator.Reduce(result, operand)
	}
	return result
}

type MathHomework []Problem

func (mh MathHomework) GrandTotal() int {
	result := 0
	for _, problem := range mh {
		result += problem.Solve()
	}
	return result
}

func InputToMathHomework(input string) MathHomework {
	lines := strings.Split(input, "\n")

	rows := make([][]int, 0)
	for _, line := range lines[:len(lines)-1] {
		numbers := strings.Fields(line)
		row := make([]int, len(numbers))
		for column, number := range numbers {
			value, _ := strconv.Atoi(number)
			row[column] = value
		}
		rows = append(rows, row)
	}

	operators := strings.Fields(lines[len(lines)-1])
	problems := make([]Problem, len(operators))
	for column, operator := range operators {
		operands := make([]int, len(rows))
		for row := range rows {
			operands[row] = rows[row][column]
		}
		problems[column] = Problem{
			Operands: operands,
			Operator: NewOperation(operator),
		}
	}
	return problems
}

func Transpose(matrix []string) []string {
	rows := len(matrix)
	cols := len(matrix[0])
	transposed := make([]string, cols)
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			transposed[i] += string(matrix[j][i])
		}
	}
	return transposed
}

func TransposeAndConvert(matrix []string) [][]int {
	var result [][]int
	var currentGroup []int
	for _, str := range Transpose(matrix) {
		trimmed := strings.TrimSpace(str)
		if trimmed == "" {
			result = append(result, currentGroup)
			currentGroup = []int{}
		} else {
			num, _ := strconv.Atoi(trimmed)
			currentGroup = append(currentGroup, num)
		}
	}
	return append(result, currentGroup)
}

func InputTransposedToMathHomework(input string) MathHomework {
	lines := strings.Split(input, "\n")
	values := TransposeAndConvert(lines[0 : len(lines)-1])
	operators := strings.Fields(lines[len(lines)-1])
	problems := make([]Problem, len(operators))
	for column, operator := range operators {
		problems[column] = Problem{
			Operands: values[column],
			Operator: NewOperation(operator),
		}
	}
	return problems
}
