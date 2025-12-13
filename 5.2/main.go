package main

import (
	"fmt"
	"math"
	"strings"

	"main/util"
)

func main() {
	blocks := util.ReadInputBlocks()
	freshRanges := blocks[0]

	// states: even indices are spoiled, odd indices are fresh
	states := []int{math.MinInt, math.MaxInt}

	for _, r := range freshRanges {
		p := strings.Split(r, "-")
		s, e := util.AssertInt(p[0]), util.AssertInt(p[1])

		states = insert(states, s, e)
	}

	util.Debugf("%v\n", states)

	cnt := 0
	for i := 1; i < len(states)-1; i += 2 {
		s, e := states[i], states[i+1]
		cnt += e - s
	}

	fmt.Printf("%d\n", cnt)
}

func insert(states []int, s int, e int) []int {
	ns := nearest(states, s, true)
	ne := nearest(states, e+1, false)

	nsFresh := ns%2 == 1
	neFresh := ne%2 == 1

	next := make([]int, ns+1)
	copy(next, states[:ns+1])

	if !nsFresh {
		if s != states[ns] {
			next = append(next, s)
		} else {
			next = next[:len(next)-1] // remove ns
		}
	}

	if neFresh {
		if e+1 != states[ne] {
			next = append(next, e+1)
		} else {
			ne += 1 // remove ne
		}
	}

	next = append(next, states[ne:]...)

	util.Debugf("Inserting %d-%d (%d, %d) into %v => %v\n", s, e, ns, ne, states, next)

	return next

	// case 1: ns is fresh and ne is spoiled, so s-e is a subrange of that => collapse the array to remove everything between ns and ne
	//
	//         0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9
	// states: 0   2       6       0     3 4   6     9
	//         ssssffffffffssssssssffffffssffffssssssf
	// s = 4, e = 11
	// ns->2, ne->13
	//         0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9
	// states: 0   2                     3 4   6     9
	//         ssssffffffffffffffffffffffssffffssssssf

	// case 2: ns is spoiled and ne is fresh, so s-e covers a brand new range => add s after ns and e+1 before ne, removing everything in between
	//
	//         0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9
	// states: 0   2       6       0     3 4   6     9
	//         ssssffffffffssssssssffffffssffffssssssf
	// s = 8, e = 17
	// ns->6, ne->19
	//	       0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9
	// states: 0   2       6   8                   8 9
	//         ssssffffffffssssffffffffffffffffffffssf

	// case 3: both ns and ne are spoiled, so s-e extends the range ended by ne => insert s after ns, remove everything between s and ne
	//
	//         0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9
	// states: 0   2       6       0     3 4   6     9
	//         ssssffffffffssssssssffffffssffffssssssf
	// s = 7, e = 15
	// ns->6, ne->16
	//         0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9
	// states: 0   2       6 7                 6     9
	//         ssssffffffffssffffffffffffffffffssssssf

	// case 4: both ns and ne are fresh, so s-e extends the range started by ns => insert e+1 before ne, remove everything between ns and e
	//
	//         0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9
	// states: 0   2       6       0     3 4   6     9
	//         ssssffffffffssssssssffffffssffffssssssf
	// s = 3, e = 17
	// ns->2, ne->19
	//         0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9
	// states: 0   2                               8 9
	//         ssssffffffffffffffffffffffffffffffffssf
}

// binary search for the nearest state, returns the index
func nearest(states []int, v int, lower bool) int {
	if len(states) == 0 {
		panic("empty states")
	}

	if len(states) == 1 {
		return 0
	}

	m := (len(states) - 1) / 2

	if states[m] == v {
		return m
	}

	if states[m] < v && states[m+1] > v {
		if lower {
			return m
		} else {
			return m + 1
		}
	}

	if states[m] < v {
		return nearest(states[m+1:], v, lower) + m + 1
	}

	return nearest(states[:m+1], v, lower)
}
