package main

import (
	"fmt"

	"main/util"
)

func main() {
	grid := util.ReadInputRuneGrid()
	next := util.ReadInputRuneGrid()

	cnt := 0

	didRemove := true
	for didRemove {
		didRemove = false

		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				next[y][x] = grid[y][x]

				if grid[y][x] != '@' {
					continue
				}

				if isAccessible(grid, x, y) {
					didRemove = true
					next[y][x] = '.'
					cnt++
				}
			}
		}

		grid, next = next, grid
	}

	fmt.Printf("%d\n", cnt)
}

func isAccessible(grid [][]rune, x, y int) bool {
	cnt := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx := x + dx
			ny := y + dy
			if ny < 0 || ny >= len(grid) || nx < 0 || nx >= len(grid[ny]) {
				continue
			}
			if grid[ny][nx] == '@' {
				cnt += 1
			}
		}
	}

	return cnt < 4
}
