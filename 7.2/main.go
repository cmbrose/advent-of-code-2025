package main

import (
	"fmt"

	"main/util"
)

func main() {
	grid := util.ReadInputRuneGrid()

	cache := make(map[int]map[int]int)

	for x := 0; x < len(grid[0]); x++ {
		if grid[0][x] == 'S' {
			splits := 1 + split(grid, x, 1, cache)
			fmt.Printf("%d\n", splits)
			return
		}
	}
}

func split(grid [][]rune, x, y int, cache map[int]map[int]int) int {
	util.Debugf("Checking %d,%d\n", x, y)
	if y >= len(grid) {
		return 0
	}

	if x < 0 || x >= len(grid[0]) {
		return 0
	}

	if grid[y][x] != '^' {
		return split(grid, x, y+1, cache)
	}

	if xcache, ok := cache[y]; ok {
		if val, ok := xcache[x]; ok {
			return val
		}
	}

	left := split(grid, x-1, y+1, cache)
	right := split(grid, x+1, y+1, cache)

	total := 1 + left + right

	if _, ok := cache[y]; !ok {
		cache[y] = make(map[int]int)
	}
	cache[y][x] = total

	return total
}
