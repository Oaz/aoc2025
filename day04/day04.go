package day04

import (
	"strings"
)

type Coordinates struct {
	X int
	Y int
}

type Map map[Coordinates]bool

func InputToMap(input string) Map {
	result := make(Map)
	for y, line := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, char := range line {
			if char == '@' {
				result[Coordinates{X: x, Y: y}] = true
			}
		}
	}
	return result
}

func Neighbors(coordinates Coordinates) []Coordinates {
	return []Coordinates{
		{X: coordinates.X - 1, Y: coordinates.Y},
		{X: coordinates.X + 1, Y: coordinates.Y},
		{X: coordinates.X, Y: coordinates.Y - 1},
		{X: coordinates.X, Y: coordinates.Y + 1},
		{X: coordinates.X - 1, Y: coordinates.Y - 1},
		{X: coordinates.X + 1, Y: coordinates.Y - 1},
		{X: coordinates.X - 1, Y: coordinates.Y + 1},
		{X: coordinates.X + 1, Y: coordinates.Y + 1},
	}
}

func CountNeighbors(m Map, c Coordinates) int {
	count := 0
	for _, neighbor := range Neighbors(c) {
		if _, ok := m[neighbor]; ok {
			count++
		}
	}
	return count
}

func CountReachableForklifts(m Map) int {
	reachable, _ := CountReachableAndGetRemainingForklifts(m)
	return reachable
}

func CountReachableAndGetRemainingForklifts(m Map) (int, Map) {
	count := 0
	remaining := make(Map)
	for coordinates := range m {
		if CountNeighbors(m, coordinates) < 4 {
			count++
		} else {
			remaining[coordinates] = true
		}
	}
	return count, remaining
}

func CountAllReachableForklifts(m Map) int {
	allReached := 0
	reached, remaining := CountReachableAndGetRemainingForklifts(m)
	for reached != 0 {
		allReached += reached
		reached, remaining = CountReachableAndGetRemainingForklifts(remaining)
	}
	return allReached
}
