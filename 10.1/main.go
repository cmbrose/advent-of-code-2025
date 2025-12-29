package main

import (
	"fmt"
	"math"
	"strings"

	"main/util"
)

func main() {
	sum := 0

	for _, line := range util.ReadInputLines() {
		sum += solve(line)
	}

	fmt.Printf("%d\n", sum)
}

func parseLine(line string) (lights int16, buttons [][]int) {
	next := func() string {
		pair := strings.SplitN(line, " ", 2)
		line = pair[len(pair)-1]
		return pair[0]
	}

	lightsRunes := []rune(next())
	lightsRunes = lightsRunes[1 : len(lightsRunes)-1]
	lights = 0
	// Reverse order so that the bitmask matches up
	for i := len(lightsRunes) - 1; i >= 0; i-- {
		lights <<= 1
		if lightsRunes[i] == '#' {
			lights |= 1
		}
	}

	for s := next(); s[0] != '{'; s = next() {
		ns := util.Map(strings.Split(s[1:len(s)-1], ","), util.AssertInt)
		buttons = append(buttons, ns)
	}

	return
}

func solve(line string) int {
	targetLights, buttons := parseLine(line)

	return solveRecursive(0, targetLights, buttons, 0)
}

func solveRecursive(currentLights, targetLights int16, buttons [][]int, index int) int {
	if currentLights == targetLights {
		return 0
	}

	if index >= len(buttons) {
		return math.MaxInt32
	}

	noPress := solveRecursive(currentLights, targetLights, buttons, index+1)

	for _, b := range buttons[index] {
		currentLights ^= 1 << b
	}

	yesPress := 1 + solveRecursive(currentLights, targetLights, buttons, index+1)

	return util.Min(noPress, yesPress)
}
