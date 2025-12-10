package main

import (
	"fmt"
	"strings"

	"main/util"
)

func main() {
	line := util.ReadInputLines()[0]

	sum := 0

	for _, r := range strings.Split(line, ",") {
		p := strings.Split(r, "-")
		s, e := util.AssertInt(p[0]), util.AssertInt(p[1])

		for i := s; i <= e; i++ {
			str := fmt.Sprintf("%d", i)
			if len(str)&1 != 0 {
				continue
			}

			mid := len(str) / 2
			first, second := str[:mid], str[mid:]
			if first == second {
				sum += i
			}
		}
	}

	fmt.Printf("%d\n", sum)
}
