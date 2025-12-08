package day08

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var debug = false

func Log(format string, a ...any) {
	if debug {
		fmt.Printf(format, a...)
	}
}

type Box struct {
	X, Y, Z int
}

func (box Box) String() string {
	return fmt.Sprintf("(%d,%d,%d)", box.X, box.Y, box.Z)
}

func parseBox(input string) Box {
	fields := strings.Split(input, ",")
	x, _ := strconv.Atoi(fields[0])
	y, _ := strconv.Atoi(fields[1])
	z, _ := strconv.Atoi(fields[2])
	return Box{X: x, Y: y, Z: z}
}

func ParseInput(input string) []Box {
	rows := strings.Split(input, "\n")
	boxes := make([]Box, len(rows))
	for row, rowContent := range rows {
		boxes[row] = parseBox(rowContent)
	}
	return boxes
}

type Connection struct {
	From, To Box
	Distance int
}

func NewConnection(from, to Box) Connection {
	dx, dy, dz := from.X-to.X, from.Y-to.Y, from.Z-to.Z
	return Connection{From: from, To: to, Distance: dx*dx + dy*dy + dz*dz}
}

func AllConnections(coordinates []Box) []Connection {
	n := len(coordinates)
	connections := make([]Connection, n*(n-1)/2)
	idx := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			connections[idx] = NewConnection(coordinates[i], coordinates[j])
			idx++
		}
	}
	sort.Slice(connections, func(i, j int) bool {
		return connections[i].Distance < connections[j].Distance
	})
	return connections
}

type Circuits struct {
	ById   map[int][]Box
	OfBox  map[Box]int
	LastId int
}

func EmptyCircuits() Circuits {
	return Circuits{ById: make(map[int][]Box), OfBox: make(map[Box]int), LastId: 0}
}

func (circuits *Circuits) GetNewCircuitId() int {
	circuits.LastId++
	return circuits.LastId
}

func (circuits *Circuits) Add(box Box, circuit int) {
	circuits.OfBox[box] = circuit
	circuits.ById[circuit] = append(circuits.ById[circuit], box)
}

func (circuits *Circuits) Merge(circuit1, circuit2 int) {
	for _, box := range circuits.ById[circuit2] {
		Log("Moving %s from circuit %d to circuit %d\n", box, circuit2, circuit1)
		circuits.Add(box, circuit1)
	}
	delete(circuits.ById, circuit2)
}

func (circuits *Circuits) AddConnection(connection Connection) {
	circuitFrom, existFrom := circuits.OfBox[connection.From]
	circuitTo, existTo := circuits.OfBox[connection.To]
	switch {
	case !existFrom && !existTo:
		circuit := circuits.GetNewCircuitId()
		Log("%s and %s create circuit %d [distance:√%d]\n", connection.From, connection.To, circuit, connection.Distance)
		circuits.Add(connection.From, circuit)
		circuits.Add(connection.To, circuit)
	case existFrom && !existTo:
		Log("%s joins circuit %d by connecting to %s [distance:√%d]\n", connection.To, circuitFrom, connection.From, connection.Distance)
		circuits.Add(connection.To, circuitFrom)
	case !existFrom:
		Log("%s joins circuit %d by connecting to %s [distance:√%d]\n", connection.From, circuitTo, connection.To, connection.Distance)
		circuits.Add(connection.From, circuitTo)
	case circuitFrom == circuitTo:
		Log("%s and %s are already connected in circuit %d : nothing happens [distance:√%d]\n", connection.From, connection.To, circuitFrom, connection.Distance)
	default:
		Log("%s is in circuit %d and %s is in circuit %d : merging [distance:√%d]\n", connection.From, circuitFrom, connection.To, circuitTo, connection.Distance)
		circuits.Merge(circuitFrom, circuitTo)
	}
}

func Part1(cs []Box, remainingConnections int) int {
	circuits := EmptyCircuits()
	for _, connection := range AllConnections(cs) {
		circuits.AddConnection(connection)
		remainingConnections--
		if remainingConnections == 0 {
			break
		}
	}
	var sizes []int
	for _, circuit := range circuits.ById {
		sizes = append(sizes, len(circuit))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	result := sizes[0] * sizes[1] * sizes[2]
	Log("%d*%d*%d = %d\n", sizes[0], sizes[1], sizes[2], result)
	return result
}

func Part2(cs []Box) int {
	circuits := EmptyCircuits()
	for _, connection := range AllConnections(cs) {
		circuits.AddConnection(connection)
		allBoxesAreInCircuit := len(circuits.OfBox) == len(cs)
		thereIsOnlyOneCircuit := len(circuits.ById) == 1
		if allBoxesAreInCircuit && thereIsOnlyOneCircuit {
			return connection.From.X * connection.To.X
		}
	}
	return -1
}
