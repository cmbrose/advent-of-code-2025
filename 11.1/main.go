package main

import (
	"fmt"
	"strings"

	"main/util"
)

func main() {
	graph := make(map[string][]string)

	for _, line := range util.ReadInputLines() {
		pair := strings.Split(line, ": ")
		node := pair[0]
		neighbors := strings.Fields(pair[1])
		graph[node] = neighbors
	}

	fmt.Printf("%d\n", countPaths(graph, "you"))
}

func countPaths(graph map[string][]string, current string) int {
	if current == "out" {
		return 1
	}

	return util.Sum(util.Map(graph[current], func(neighbor string) int {
		return countPaths(graph, neighbor)
	}))
}
