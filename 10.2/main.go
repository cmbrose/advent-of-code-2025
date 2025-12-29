package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"main/util"
)

func main() {
	sum := 0

	for _, line := range util.ReadInputLines() {
		pushes := solve(line)
		sum += pushes
		util.Debugf("%s -> %d\n", line, pushes)
	}

	fmt.Printf("%d\n", sum)
}

func parseLine(line string) (buttons [][]int, targetCounter []int) {
	next := func() string {
		pair := strings.SplitN(line, " ", 2)
		line = pair[len(pair)-1]
		return pair[0]
	}

	// skip lights
	next()

	for {
		s := next()
		ns := util.Map(strings.Split(s[1:len(s)-1], ","), util.AssertInt)

		if s[0] == '{' {
			targetCounter = ns
			return
		}

		buttons = append(buttons, ns)
	}
}

func solve(line string) int {
	buttons, targetCounter := parseLine(line)

	return solveRecursive(make([]int, len(targetCounter)), targetCounter, buttons, 0, make(map[cacheKey]int))
}

type cacheKey struct {
	current string
	index   int
}

func solveRecursive(currentCounter, targetCounter []int, buttons [][]int, index int, cache map[cacheKey]int) int {
	if slices.Equal(currentCounter, targetCounter) {
		return 0
	}

	key := cacheKey{fmt.Sprintf("%v", currentCounter), index}
	if val, ok := cache[key]; ok {
		return val
	}

	if index >= len(buttons) || !isValid(currentCounter, targetCounter, buttons[index:]) {
		return math.MaxInt32
	}

	noPress := solveRecursive(currentCounter, targetCounter, buttons, index+1, cache)

	for _, n := range buttons[index] {
		currentCounter[n] += 1
	}
	yesPress := 1 + solveRecursive(currentCounter, targetCounter, buttons, index, cache)

	// "unpress"
	for _, n := range buttons[index] {
		currentCounter[n] -= 1
	}

	best := util.Min(noPress, yesPress)
	cache[key] = best
	return best
}

func isValid(counter, targetCounter []int, buttons [][]int) bool {
	for i := range counter {
		if counter[i] > targetCounter[i] {
			return false
		}
	}

	return true
}
