package main

import (
	"fmt"

	"main/util"
)

func main() {
	sum := 0

	for _, row := range util.ParseIntGrid() {
		arr := make([]int, 12)

		for i, cell := range row {
			for j := 0; j < 12; j++ {
				if cell > arr[j] && i < len(row)-11+j {
					arr[j] = cell
					for k := j + 1; k < 12; k++ {
						arr[k] = 0
					}
					break
				}
			}
		}

		util.Debugf("%v => %v\n", row, arr)

		val := 0
		for _, v := range arr {
			val = val*10 + v
		}
		util.Debugf("value: %d\n", val)

		sum += val
	}

	fmt.Printf("%d\n", sum)
}
