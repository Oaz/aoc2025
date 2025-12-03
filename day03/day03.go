package day03

import (
	"iter"
	"strconv"
	"strings"

	"github.com/alvii147/gloop"
)

type Battery struct {
	Index int
	Value int
}
type Bank []Battery

func isBeforeInJoltageOrder(a, b Battery) bool {
	if a.Value == b.Value {
		return a.Index < b.Index
	}
	return a.Value > b.Value
}

func hasIndexGreaterThan(n int) gloop.FilterFunc[Battery] {
	return func(battery Battery) bool { return battery.Index > n }
}

func hasIndexDifferentThan(n int) gloop.FilterFunc[Battery] {
	return func(battery Battery) bool { return battery.Index != n }
}

func InputToBanks(input string) []Bank {
	banks := make([]Bank, 0)
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		bank := make(Bank, 0)
		for index, char := range line {
			value, _ := strconv.Atoi(string(char))
			bank = append(bank, Battery{Index: index, Value: value})
		}
		banks = append(banks, bank)
	}
	return banks
}

func (bank Bank) Digits(numberOfBatteriesToTurnOn int) iter.Seq[int] {
	return func(yield func(int) bool) {
		available := gloop.Slice(bank)
		reserve := gloop.Collect[Battery]()
		for numberOfBatteriesToTurnOn > 0 {
			maxIndex := len(bank) - numberOfBatteriesToTurnOn + 1
			maxValue := gloop.MinByComparison(available, isBeforeInJoltageOrder)
			if maxValue.Index < maxIndex {
				if !yield(maxValue.Value) {
					return
				}
				numberOfBatteriesToTurnOn--
				available = gloop.Chain(gloop.Filter(available, hasIndexGreaterThan(maxValue.Index)), reserve)
				reserve = gloop.Collect[Battery]()
			} else {
				available = gloop.Filter(available, hasIndexDifferentThan(maxValue.Index))
				reserve = gloop.Chain(reserve, gloop.Collect(maxValue))
			}
		}
	}
}

func (bank Bank) Joltage(numberOfBatteriesToTurnOn int) int {
	return gloop.Fold(bank.Digits(numberOfBatteriesToTurnOn), func(result, digit int) int {
		return result*10 + digit
	})
}

func TotalOutputJoltage(banks []Bank, n int) int {
	return gloop.Sum(gloop.Transform(gloop.Slice(banks), func(bank Bank) int {
		return bank.Joltage(n)
	}))
}
