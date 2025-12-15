package main

import (
	"fmt"
	"strings"

	"main/util"
)

func main() {
	lines := util.ReadInputLines()
	ops := strings.Fields(lines[len(lines)-1])
	vals := util.Map(lines[:len(lines)-1], func(s string) []int {
		return util.Map(strings.Fields(s), util.AssertInt)
	})

	sum := 0
	for i, op := range ops {
		switch op {
		case "+":
			sum += util.Sum(util.Map(vals, func(v []int) int { return v[i] }))
		case "*":
			sum += util.Product(util.Map(vals, func(v []int) int { return v[i] }))
		}
	}

	fmt.Printf("%d\n", sum)
}
