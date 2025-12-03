package day03

import (
	"sort"
	"strconv"
	"strings"
)

type Battery struct {
	Index int
	Value int
}
type Bank []Battery

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

func (bank Bank) Joltage(numberOfBatteriesToTurnOn int) int {
	available := make([]Battery, len(bank))
	copy(available, bank)
	reserve := make([]Battery, 0)
	result := 0
	for numberOfBatteriesToTurnOn > 0 {
		sort.Slice(available, func(i, j int) bool {
			if available[i].Value == available[j].Value {
				return available[i].Index < available[j].Index
			}
			return available[i].Value > available[j].Value
		})
		maxValue := available[0]
		if maxValue.Index < len(bank)-numberOfBatteriesToTurnOn+1 {
			result = result*10 + maxValue.Value
			numberOfBatteriesToTurnOn--
			remainOnRightOfTurnedOn := make([]Battery, 0)
			for _, battery := range available[1:] {
				if battery.Index > maxValue.Index {
					remainOnRightOfTurnedOn = append(remainOnRightOfTurnedOn, battery)
				}
			}
			for _, battery := range reserve {
				remainOnRightOfTurnedOn = append(remainOnRightOfTurnedOn, battery)
			}
			available = remainOnRightOfTurnedOn
			reserve = make([]Battery, 0)
		} else {
			reserve = append(reserve, maxValue)
			available = available[1:]
		}
	}
	return result
}

func TotalOutputJoltage(banks []Bank, n int) int {
	total := 0
	for _, bank := range banks {
		total += bank.Joltage(n)
	}
	return total
}
