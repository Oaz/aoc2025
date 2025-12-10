package day09

import (
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	X, Y int
}

func (t Pair) String() string {
	return fmt.Sprintf("(%d,%d)", t.X, t.Y)
}

func parseTile(input string) Pair {
	fields := strings.Split(input, ",")
	x, _ := strconv.Atoi(fields[0])
	y, _ := strconv.Atoi(fields[1])
	return Pair{X: x, Y: y}
}

type Interval struct {
	Min, Max int
}

func (i Interval) Length() int {
	return i.Max - i.Min + 1
}

type XSegment struct {
	XInterval Interval
	Y         int
}

func createXSegment(tile Pair, previousTile Pair) XSegment {
	xSegment := XSegment{
		XInterval: Interval{
			Min: min(previousTile.X, tile.X),
			Max: max(previousTile.X, tile.X),
		},
		Y: tile.Y,
	}
	return xSegment
}

type YSegment struct {
	X         int
	YInterval Interval
}

func createYSegment(tile Pair, previousTile Pair) YSegment {
	ySegment := YSegment{
		X: tile.X,
		YInterval: Interval{
			Min: min(previousTile.Y, tile.Y),
			Max: max(previousTile.Y, tile.Y),
		},
	}
	return ySegment
}

type Floor struct {
	RedTiles    []Pair
	HasRedTiles map[Pair]bool
	XSegments   []XSegment
	YSegments   []YSegment
}

func (floor *Floor) NextTile(tile Pair, addTile bool) {
	previousTile := floor.RedTiles[len(floor.RedTiles)-1]
	if addTile {
		floor.RedTiles = append(floor.RedTiles, tile)
		floor.HasRedTiles[tile] = true
	}
	if tile.Y == previousTile.Y {
		floor.XSegments = append(floor.XSegments, createXSegment(tile, previousTile))
	} else {
		floor.YSegments = append(floor.YSegments, createYSegment(tile, previousTile))
	}
}

func (floor *Floor) HasRedCorners(xs Interval, ys Interval) bool {
	topLeft := floor.HasRedTiles[Pair{X: xs.Min, Y: ys.Min}]
	topRight := floor.HasRedTiles[Pair{X: xs.Max, Y: ys.Min}]
	bottomLeft := floor.HasRedTiles[Pair{X: xs.Min, Y: ys.Max}]
	bottomRight := floor.HasRedTiles[Pair{X: xs.Max, Y: ys.Max}]
	return (topLeft && bottomRight) || (topRight && bottomLeft)
}

func ParseInput(input string) Floor {
	rows := strings.Split(input, "\n")
	floor := Floor{
		RedTiles:    make([]Pair, 0),
		HasRedTiles: map[Pair]bool{},
		XSegments:   make([]XSegment, 0),
		YSegments:   make([]YSegment, 0),
	}
	for row, rowContent := range rows {
		tile := parseTile(rowContent)
		if row == 0 {
			floor.RedTiles = append(floor.RedTiles, tile)
			continue
		}
		floor.NextTile(tile, true)
	}
	floor.NextTile(floor.RedTiles[0], false)
	return floor
}

func Part1(floor Floor) int {
	maxArea := 0
	for i := 0; i < len(floor.RedTiles); i++ {
		tile1 := floor.RedTiles[i]
		for j := i + 1; j < len(floor.RedTiles); j++ {
			tile2 := floor.RedTiles[j]
			minX := min(tile1.X, tile2.X)
			maxX := max(tile1.X, tile2.X)
			minY := min(tile1.Y, tile2.Y)
			maxY := max(tile1.Y, tile2.Y)
			width := maxX - minX + 1
			height := maxY - minY + 1
			area := width * height
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

type Zones struct {
	Xs, Ys []int
}

type PointFinder struct {
	XIntervals map[int][]Interval
}

func (floor *Floor) Analyze() (Zones, PointFinder) {
	zones := Zones{Xs: make([]int, 0), Ys: make([]int, 0)}
	pointFinder := PointFinder{XIntervals: make(map[int][]Interval)}
	for _, xSegment := range floor.XSegments {
		if _, exist := pointFinder.XIntervals[xSegment.Y]; !exist {
			zones.Ys = append(zones.Ys, xSegment.Y)
			pointFinder.XIntervals[xSegment.Y] = make([]Interval, 0)
		}
		pointFinder.XIntervals[xSegment.Y] = append(pointFinder.XIntervals[xSegment.Y], xSegment.XInterval)
	}
	sort.Ints(zones.Ys)
	xs := make(map[int]bool)
	for _, ySegment := range floor.YSegments {
		if _, exist := xs[ySegment.X]; exist {
			continue
		}
		zones.Xs = append(zones.Xs, ySegment.X)
	}
	sort.Ints(zones.Xs)
	return zones, pointFinder
}

func (pf PointFinder) FindPointInHorizontalSegments(x, y int) bool {
	for _, interval := range pf.XIntervals[y] {
		if x >= interval.Min && x <= interval.Max {
			return true
		}
	}
	return false
}

func (zones Zones) FindGreen(pf PointFinder) []big.Int {
	greens := make([]big.Int, len(zones.Xs))
	for i := 0; i < len(zones.Xs)-1; i++ {
		x := zones.Xs[i] + 1
		inside := false
		for j := 0; j < len(zones.Ys); j++ {
			if pf.FindPointInHorizontalSegments(x, zones.Ys[j]) {
				inside = !inside
			}
			if inside {
				greens[i].SetBit(&greens[i], j, 1)
			}
		}
	}
	return greens
}

func longestBitSequence(x *big.Int) (int, int) {
	maxLen := 0
	endMaxLen := 0
	currentLen := 0
	n := new(big.Int).Set(x)
	for i := 0; i < n.BitLen(); i++ {
		if n.Bit(i) == 1 {
			currentLen++
			if currentLen > maxLen {
				maxLen = currentLen
				endMaxLen = i
			}
		} else {
			currentLen = 0
		}
	}
	return endMaxLen - maxLen + 1, endMaxLen
}

func FindLongestVerticalGreenArea(greens []big.Int, iStart int, iEnd int) (int, int) {
	green := greens[iStart]
	for i := iStart + 1; i <= iEnd; i++ {
		var result big.Int
		result.And(&green, &greens[i])
		green = result
	}
	return longestBitSequence(&green)
}

func Part2(floor Floor) int {
	zones, pointFinder := floor.Analyze()
	greens := zones.FindGreen(pointFinder)
	maxArea := 0
	for iStart := 0; iStart < len(zones.Xs)-1; iStart++ {
		for iEnd := iStart; iEnd < len(zones.Xs)-1; iEnd++ {
			jStart, jEnd := FindLongestVerticalGreenArea(greens, iStart, iEnd)
			xs := Interval{Min: zones.Xs[iStart], Max: zones.Xs[iEnd+1]}
			ys := Interval{Min: zones.Ys[jStart], Max: zones.Ys[jEnd+1]}
			if floor.HasRedCorners(xs, ys) {
				area := xs.Length() * ys.Length()
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}
