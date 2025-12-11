package day11

import (
	"fmt"
	"os"
	"strings"
)

type Graph struct {
	SortedNodes []string
	Edges       map[string][]string
}

func ParseInput(input string) Graph {
	rows := strings.Split(input, "\n")
	edges := make(map[string][]string)
	for _, rowContent := range rows {
		parts := strings.Split(rowContent, " ")
		parts[0] = parts[0][:len(parts[0])-1]
		for _, neighbor := range parts[1:] {
			edges[parts[0]] = append(edges[parts[0]], neighbor)
		}
	}
	return Graph{SortedNodes: Sort(edges), Edges: edges}
}

func Sort(edges map[string][]string) []string {
	degree := make(map[string]int)
	for node, neighbors := range edges {
		if _, exists := degree[node]; !exists {
			degree[node] = 0
		}
		for _, neighbor := range neighbors {
			degree[neighbor]++
		}
	}
	var queue []string
	for node, deg := range degree {
		if deg == 0 {
			queue = append(queue, node)
		}
	}
	sorted := make([]string, 0)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		sorted = append(sorted, node)
		for _, neighbor := range edges[node] {
			degree[neighbor]--
			if degree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	return sorted
}

func (graph Graph) CountPaths(start string, end string) int {
	counts := make(map[string]int)
	counts[start] = 1
	for _, node := range graph.SortedNodes {
		if node == end {
			break
		}
		for _, neighbor := range graph.Edges[node] {
			counts[neighbor] += counts[node]
		}
	}
	return counts[end]
}

func Part1(graph Graph) int {
	return graph.CountPaths("you", "out")
}

func Part2(graph Graph) int {
	svrToFft := graph.CountPaths("svr", "fft")
	fftToDac := graph.CountPaths("fft", "dac")
	dacToOut := graph.CountPaths("dac", "out")
	return svrToFft * fftToDac * dacToOut
}

func CreateDotFile(graph Graph) {
	file, _ := os.Create("graph.dot")
	defer file.Close()
	file.WriteString("digraph {\n")
	file.WriteString("\tyou [style=filled, fillcolor=red]\n")
	file.WriteString("\tout [style=filled, fillcolor=red]\n")
	file.WriteString("\tsvr [style=filled, fillcolor=red]\n")
	file.WriteString("\tfft [style=filled, fillcolor=red]\n")
	file.WriteString("\tdac [style=filled, fillcolor=red]\n")
	for node, neighbors := range graph.Edges {
		for _, neighbor := range neighbors {
			file.WriteString(fmt.Sprintf("\t%s -> %s\n", node, neighbor))
		}
	}
	file.WriteString("}\n")
}
