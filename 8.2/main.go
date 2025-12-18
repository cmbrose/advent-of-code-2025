package main

import (
	"fmt"
	"strings"

	"main/util"
)

type Breaker struct {
	x, y, z int
	circuit *Circuit
	id      int
}

type Circuit struct {
	breakers []*Breaker
	id       int
}

type Connection struct {
	b1, b2 *Breaker
	distSq int
}

func main() {
	circuits := make(map[*Circuit]struct{})

	id := 0
	breakers := util.MapInputLines(func(line string) *Breaker {
		vals := strings.Split(line, ",")
		b := Breaker{
			id:      id,
			x:       util.AssertInt(vals[0]),
			y:       util.AssertInt(vals[1]),
			z:       util.AssertInt(vals[2]),
			circuit: &Circuit{id: id},
		}
		id += 1

		circuits[b.circuit] = struct{}{}
		b.circuit.breakers = append(b.circuit.breakers, &b)

		util.Debugf("Breaker %d: (%d, %d, %d)\n", b.id, b.x, b.y, b.z)

		return &b
	})

	queue := util.NewPriorityQueue(func(c1, c2 Connection) bool {
		return c1.distSq < c2.distSq
	})

	for i, b1 := range breakers[:len(breakers)-1] {
		for _, b2 := range breakers[i+1:] {
			distSq := (b1.x - b2.x) * (b1.x - b2.x)
			distSq += (b1.y - b2.y) * (b1.y - b2.y)
			distSq += (b1.z - b2.z) * (b1.z - b2.z)

			queue.Push(Connection{
				b1:     b1,
				b2:     b2,
				distSq: distSq,
			})
		}
	}

	c := Connection{}
	for len(circuits) > 1 {
		c = queue.Pop()

		util.Debugf("Checking %d:(%d,%d,%d) -> %d:(%d,%d,%d) [%d]\n", c.b1.id, c.b1.x, c.b1.y, c.b1.z, c.b2.id, c.b2.x, c.b2.y, c.b2.z, c.distSq)

		if c.b1.circuit != c.b2.circuit {
			// merge circuits
			oldCircuit := c.b2.circuit
			newCircuit := c.b1.circuit

			for _, b := range oldCircuit.breakers {
				b.circuit = newCircuit
			}
			delete(circuits, oldCircuit)
			util.Debugf("Merging circuits id %d,%d with %d and %d breakers, total of %d circuits\n", newCircuit.id, oldCircuit.id, len(newCircuit.breakers), len(oldCircuit.breakers), len(circuits))

			newCircuit.breakers = append(newCircuit.breakers, oldCircuit.breakers...)
		} else {
			util.Debugf("Skipping connection, same circuit id %d\n", c.b1.circuit.id)
		}
	}

	fmt.Printf("%d\n", c.b1.x*c.b2.x)
}
