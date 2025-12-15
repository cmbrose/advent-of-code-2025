package main

import (
	"fmt"

	"main/util"
)

func main() {
	lines := util.ReadInputLines()
	ops := lines[len(lines)-1]
	vals := lines[:len(lines)-1]

	sum := 0
	for ops != "" {
		op := ops[0]

		ops = ops[1:]
		offset := 1
		for ops != "" && ops[0] == ' ' {
			offset += 1
			ops = ops[1:]
		}

		if ops == "" {
			offset += 1
		}

		nums := make([]string, len(vals))
		for i, v := range vals {
			nums[i] = v[:offset-1] // skip the blank column

			vals[i] = v
			if ops != "" {
				vals[i] = v[offset:]
			}
		}

		trueNums := make([]int, offset-1)
		for i := 0; i < offset-1; i += 1 {
			for _, n := range nums {
				if n[i] != ' ' {
					trueNums[i] = trueNums[i]*10 + int(n[i]-'0')
				}
			}
		}

		util.Debugf("Evaluating (%c) on %v\n", op, trueNums)

		switch op {
		case '+':
			sum += util.Sum(trueNums)
		case '*':
			sum += util.Product(trueNums)
		}
	}

	fmt.Printf("%d\n", sum)
}
