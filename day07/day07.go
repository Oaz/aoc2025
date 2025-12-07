package day07

import (
	"strings"
)

type Coordinates struct {
	X int
	Y int
}

func (c Coordinates) Add(other Coordinates) Coordinates {
	return Coordinates{
		X: c.X + other.X,
		Y: c.Y + other.Y,
	}
}

type Manifold struct {
	Start     Coordinates
	Height    int
	Splitters map[Coordinates]bool
}

func InputToManifold(input string) Manifold {
	rows := strings.Split(input, "\n")
	manifold := Manifold{
		Start:     Coordinates{},
		Height:    len(rows),
		Splitters: make(map[Coordinates]bool),
	}
	for row, rowContent := range rows {
		for col, char := range rowContent {
			coordinates := Coordinates{X: col, Y: row}
			switch char {
			case '^':
				manifold.Splitters[coordinates] = true
			case 'S':
				manifold.Start = coordinates
			}
		}
	}
	return manifold
}

type BeamSequence struct {
	Beams     map[Coordinates]int
	Splits    int
	Timelines int
}

func (manifold Manifold) InitialBeamSequence() BeamSequence {
	beams := make(map[Coordinates]int)
	beams[manifold.Start] = 1
	return BeamSequence{
		Beams:     beams,
		Splits:    0,
		Timelines: 0,
	}
}

func (manifold Manifold) MoveAndSplitBeams(s BeamSequence) BeamSequence {
	newBeams := make(map[Coordinates]int)
	for coordinates, count := range s.Beams {
		coordinates = coordinates.Add(Coordinates{X: 0, Y: 1})
		if coordinates.Y > manifold.Height {
			s.Timelines += count
			continue
		}
		if manifold.Splitters[coordinates] {
			s.Splits++
			newBeams[coordinates.Add(Coordinates{X: -1, Y: 0})] += count
			newBeams[coordinates.Add(Coordinates{X: 1, Y: 0})] += count
		} else {
			newBeams[coordinates] += count
		}
	}
	return BeamSequence{
		Beams:     newBeams,
		Timelines: s.Timelines,
		Splits:    s.Splits,
	}
}
func (manifold Manifold) Beam() BeamSequence {
	sequence := manifold.InitialBeamSequence()
	for len(sequence.Beams) > 0 {
		sequence = manifold.MoveAndSplitBeams(sequence)
	}
	return sequence
}

func (manifold Manifold) CountSplits() int { return manifold.Beam().Splits }

func (manifold Manifold) CountTimelines() int { return manifold.Beam().Timelines }
