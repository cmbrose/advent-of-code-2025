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

	fmt.Printf("%d\n", countPaths(graph, "svr"))
}

func countPaths(graph map[string][]string, current string) int {
	return countPathsRec(graph, current, false, false, make(map[string]int64))
}

func countPathsRec(graph map[string][]string, current string, hasDac, hasFft bool, cache map[string]int64) int {
	if current == "out" {
		if hasDac && hasFft {
			return 1
		}
		return 0
	}

	if val, found := cache[fmt.Sprintf("%s-%t-%t", current, hasDac, hasFft)]; found {
		return int(val)
	}

	if current == "dac" {
		hasDac = true
	}
	if current == "fft" {
		hasFft = true
	}

	val := util.Sum(util.Map(graph[current], func(neighbor string) int {
		return countPathsRec(graph, neighbor, hasDac, hasFft, cache)
	}))
	cache[fmt.Sprintf("%s-%t-%t", current, hasDac, hasFft)] = int64(val)
	return val
}
