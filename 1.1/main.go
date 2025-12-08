package main

import (
	"fmt"

	"main/util"
)

func main() {
	pos := 50
	zeros := 0

	for _, line := range util.ReadInputLines() {
		dir := 1
		if line[0] == 'L' {
			dir = -1
		}

		mag := util.AssertInt(line[1:])

		oldPos := pos
		pos = (pos + dir*mag + 100) % 100

		if pos == 0 {
			zeros += 1
		}

		fmt.Printf("Rotate %s from %d, landed on %d => %d\n", line, oldPos, pos, zeros)
	}

	fmt.Printf("%d\n", zeros)
}
