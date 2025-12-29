package main

import (
	"fmt"
	"strings"

	"main/util"
)

type Point struct {
	x, y int
}

func main() {
	points := util.MapInputLines(func(line string) Point {
		p := strings.Split(line, ",")
		return Point{
			x: util.AssertInt(p[0]),
			y: util.AssertInt(p[1]),
		}
	})

	maxArea := 0
	for i, p1 := range points[:len(points)-1] {
		for _, p2 := range points[i+1:] {
			area := (util.Abs(p1.x-p2.x) + 1) * (util.Abs(p1.y-p2.y) + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	fmt.Printf("%d\n", maxArea)
}
