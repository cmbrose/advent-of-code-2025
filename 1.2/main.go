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

		fullMag := util.AssertInt(line[1:])
		mag := fullMag % 100
		zeros += fullMag / 100

		oldPos := pos
		testPos := (pos + dir*mag)
		pos = (testPos + 100) % 100

		if (oldPos != 0 && testPos <= 0) || testPos >= 100 {
			zeros += 1
		}

		fmt.Printf("Rotate %s from %d, landed on %d => %d\n", line, oldPos, pos, zeros)
	}

	fmt.Printf("%d\n", zeros)
}
