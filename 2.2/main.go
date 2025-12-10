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

			for j := 0; j < len(str)/2+1; j++ {
				left, right := str[:j], str[j:]

				repCnt := strings.Count(right, left)
				if repCnt*len(left) == len(right) {
					sum += i
					break
				}
			}
		}
	}

	fmt.Printf("%d\n", sum)
}
