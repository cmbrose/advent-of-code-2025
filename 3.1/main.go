package main

import (
	"fmt"

	"main/util"
)

func main() {
	sum := 0

	for _, row := range util.ParseIntGrid() {
		d1 := row[0]
		d2 := 0
		for i, cell := range row[1:] {
			if cell > d1 && i < len(row)-2 {
				d1 = cell
				d2 = 0
			} else if cell > d2 {
				d2 = cell
			}
		}

		util.Debugf("%v => %d, %d\n", row, d1, d2)

		sum += d1*10 + d2
	}

	fmt.Printf("%d\n", sum)
}
