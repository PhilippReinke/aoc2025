package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("error reading from stdin:", err)
		os.Exit(1)
	}

	dataTrimmed := strings.TrimRight(string(data), "\n")
	lines := strings.Split(string(dataTrimmed), "\n")

	// parse input
	var coordinates []Coordinate
	for _, line := range lines {
		numsString := strings.Split(line, ",")
		if len(numsString) != 3 {
			panic("unexpected count of nums")
		}

		x, err := strconv.Atoi(numsString[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(numsString[1])
		if err != nil {
			panic(err)
		}
		z, err := strconv.Atoi(numsString[2])
		if err != nil {
			panic(err)
		}

		coordinates = append(coordinates, NewCoord(x, y, z))
	}

	// find nearest pairs
	var pairs []Pair
	for i := range coordinates {
		for j := i + 1; j < len(coordinates); j++ {
			pairs = append(pairs, Pair{
				box1: coordinates[i],
				box2: coordinates[j],
				dist: coordinates[i].Distance(coordinates[j]),
			})
		}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].dist < pairs[j].dist
	})

	// exa considers 10 examples
	// real input considers 1000 examples
	numPairs := 1000
	if len(pairs) < numPairs {
		numPairs = 10
	}

	// connect them
	var circuits Circuits
	for _, pair := range pairs[:numPairs] {
		circuits.AddPair(pair.box1, pair.box2)
	}
	circuits.Merge()
	l1, l2, l3 := circuits.Longest()

	// part 2
	var circuits2 Circuits
	for _, coord := range coordinates {
		circuits2 = append(circuits2, Circuit{coord: struct{}{}})
	}
	var sol2 int
	for _, pair := range pairs {
		circuits2.AddPair(pair.box1, pair.box2)
		circuits2.Merge()
		if len(circuits2) == 1 {
			sol2 = pair.box1.x * pair.box2.x
			break
		}
	}

	fmt.Println("sol1:", l1*l2*l3)
	fmt.Println("sol2:", sol2)
}

type Pair struct {
	box1 Coordinate
	box2 Coordinate
	dist float64
}

type Coordinate struct {
	x, y, z int
}

func NewCoord(x, y, z int) Coordinate {
	return Coordinate{x, y, z}
}

func (c Coordinate) Distance(other Coordinate) float64 {
	xDist := float64(c.x - other.x)
	yDist := float64(c.y - other.y)
	zDist := float64(c.z - other.z)
	return math.Sqrt(xDist*xDist + yDist*yDist + zDist*zDist)
}

type Circuits []Circuit
type Circuit map[Coordinate]struct{}

func (c *Circuits) AddPair(a, b Coordinate) {
	for _, circuit := range *c {
		_, okA := circuit[a]
		_, okB := circuit[b]
		if okA || okB {
			circuit[a] = struct{}{}
			circuit[b] = struct{}{}
			return
		}
	}

	// no circuit has a or b
	*c = append(*c, Circuit{
		a: struct{}{},
		b: struct{}{},
	})
}

func (c *Circuits) Longest() (int, int, int) {
	var lengths []int
	for _, circuit := range *c {
		lengths = append(lengths, len(circuit))
	}
	sort.Slice(lengths, func(i, j int) bool {
		return lengths[i] > lengths[j]
	})
	return lengths[0], lengths[1], lengths[2]
}

func (c *Circuits) Merge() {
	merged := true
	for merged {
		merged = false
		for i := 0; i < len(*c); i++ {
			for j := i + 1; j < len(*c); j++ {
				if circuitsOverlap((*c)[i], (*c)[j]) {
					for coord := range (*c)[j] {
						(*c)[i][coord] = struct{}{}
					}

					*c = append((*c)[:j], (*c)[j+1:]...)
					merged = true
					break
				}
			}
			if merged {
				break
			}
		}
	}
}

func circuitsOverlap(a, b Circuit) bool {
	for coord := range a {
		if _, ok := b[coord]; ok {
			return true
		}
	}
	return false
}
