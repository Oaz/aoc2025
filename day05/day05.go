package day05

import (
	"sort"
	"strconv"
	"strings"
)

type FreshRange struct {
	Low, High int
}

func (r FreshRange) Size() int {
	return r.High - r.Low + 1
}

func (r FreshRange) Contains(value int) bool {
	return value >= r.Low && value <= r.High
}

type Inventory struct {
	Fresh     []FreshRange
	Available []int
}

func InputToInventory(input string) Inventory {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	fresh := make([]FreshRange, 0)
	for _, r := range strings.Split(strings.TrimSpace(parts[0]), "\n") {
		bounds := strings.Split(r, "-")
		low, _ := strconv.Atoi(bounds[0])
		high, _ := strconv.Atoi(bounds[1])
		fresh = append(fresh, FreshRange{Low: low, High: high})
	}
	items := make([]int, 0)
	for _, item := range strings.Split(strings.TrimSpace(parts[1]), "\n") {
		value, _ := strconv.Atoi(item)
		items = append(items, value)
	}
	return Inventory{
		Fresh:     mergeRanges(fresh),
		Available: items,
	}
}

func mergeRanges(ranges []FreshRange) []FreshRange {
	sort.Slice(ranges, func(i, j int) bool { return ranges[i].Low < ranges[j].Low })
	result := make([]FreshRange, 0)
	for _, r := range ranges {
		if len(result) == 0 || result[len(result)-1].High < r.Low-1 {
			result = append(result, r)
		} else {
			result[len(result)-1].High = max(r.High, result[len(result)-1].High)
		}
	}
	return result
}

func (inventory Inventory) isFresh(ingredient int) bool {
	for _, r := range inventory.Fresh {
		if r.Contains(ingredient) {
			return true
		}
	}
	return false
}

func (inventory Inventory) CountFreshAvailableIngredients() int {
	count := 0
	for _, item := range inventory.Available {
		if inventory.isFresh(item) {
			count++
		}
	}
	return count
}

func (inventory Inventory) CountFreshIngredients() int {
	count := 0
	for _, r := range inventory.Fresh {
		count += r.Size()
	}
	return count
}
