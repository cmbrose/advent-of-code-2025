package main

import (
	"fmt"

	"main/util"
)

func main() {
	grid := util.ReadInputRuneGrid()

	beams := make(map[int]bool)

	for x := 0; x < len(grid[0]); x++ {
		if grid[0][x] == 'S' {
			beams[x] = true
			break
		}
	}

	splits := 0

	for y := 1; y < len(grid); y += 1 {
		for x := range beams {
			if grid[y][x] == '^' {
				delete(beams, x)
				splits += 1

				if _, ok := beams[x-1]; !ok && x-1 >= 0 {
					beams[x-1] = true
				}
				if _, ok := beams[x+1]; !ok && x+1 < len(grid[0]) {
					beams[x+1] = true
				}
			}
		}

		for x := 0; x < len(grid[0]); x += 1 {
			if _, ok := beams[x]; ok {
				util.Debugf("|")
			} else {
				util.Debugf(string(grid[y][x]))
			}
		}
		util.Debugf(" %d \n", splits)
	}

	fmt.Printf("%d\n", splits)
}
