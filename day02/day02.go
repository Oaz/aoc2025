package day02

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Min       int
	MinDigits int
	Max       int
	MaxDigits int
}

func getBound(input string) int {
	value, _ := strconv.Atoi(input)
	return value
}

func InputToListOfRanges(input string) []Range {
	ranges := make([]Range, 0)
	for _, pair := range strings.Split(input, ",") {
		split := strings.Split(pair, "-")
		minValue := getBound(split[0])
		maxValue := getBound(split[1])
		ranges = append(ranges, Range{
			Min:       minValue,
			MinDigits: numberOfDigits(minValue),
			Max:       maxValue,
			MaxDigits: numberOfDigits(maxValue),
		})
	}
	return ranges
}

func numberOfDigits(n int) int {
	return int(math.Log10(float64(n))) + 1
}

func pow10(n int) int {
	return int(math.Pow(10, float64(n)))
}

func fakeID(n, count int) int {
	factor := pow10(numberOfDigits(n))
	result := 0
	for i := 0; i < count; i++ {
		result = result*factor + n
	}
	return result
}

func getFakeIDs(r Range, count int) []int {
	var result []int
	for digits := r.MinDigits; digits <= r.MaxDigits; digits++ {
		minFakeCandidate := pow10(digits - 1)
		if digits == r.MinDigits {
			minFakeCandidate = r.Min
		}
		maxFakeCandidate := pow10(digits) - 1
		if digits == r.MaxDigits {
			maxFakeCandidate = r.Max
		}
		seed := pow10((digits+count-1)/count - 1)
		for ; ; seed++ {
			id := fakeID(seed, count)
			if id > maxFakeCandidate {
				break
			}
			if id >= minFakeCandidate {
				result = append(result, id)
			}
		}
	}
	return result
}

type FakeIdPolicy struct {
	MinRepeat int
	MaxRepeat int
}

func RepeatTwicePolicy() FakeIdPolicy {
	return FakeIdPolicy{MinRepeat: 2, MaxRepeat: 2}
}

func RepeatAtLeastTwicePolicy() FakeIdPolicy {
	return FakeIdPolicy{MinRepeat: 2, MaxRepeat: 10}
}

func (p FakeIdPolicy) FindFakeIDs(ranges []Range) []int {
	allIDs := make([]int, 0)
	seen := make(map[int]bool)
	for _, r := range ranges {
		for count := p.MinRepeat; count <= p.MaxRepeat; count++ {
			ids := getFakeIDs(r, count)
			for _, id := range ids {
				if !seen[id] {
					seen[id] = true
					allIDs = append(allIDs, id)
				}
			}
		}
	}
	return allIDs
}

func FindFakeIDs(r Range, policy FakeIdPolicy) []int {
	result := policy.FindFakeIDs([]Range{r})
	sort.Ints(result)
	return result
}

func SumFakeIDs(ranges []Range, policy FakeIdPolicy) int {
	result := 0
	for _, fakeId := range policy.FindFakeIDs(ranges) {
		result += fakeId
	}
	return result
}
