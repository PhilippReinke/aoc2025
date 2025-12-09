package main

import (
	"fmt"
	"io"
	"os"
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

	var coords []Coord
	for _, line := range lines {
		split := strings.Split(line, ",")
		if len(split) != 2 {
			panic("expected 2")
		}

		x, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		coords = append(coords, NewCoord(x, y))
	}

	// prep for part 2
	boundary := map[Coord]struct{}{}
	for i := range len(coords) {
		j := (i + 1) % len(coords)

		for _, coord := range CoordsBetween(coords[i], coords[j]) {
			boundary[coord] = struct{}{}
		}
	}

	var sol1, sol2 int
	for i := range coords {
		for j := i + 1; j < len(coords); j++ {
			c1 := coords[i]
			c2 := coords[j]

			width := abs(c1.x-c2.x) + 1
			height := abs(c1.y-c2.y) + 1
			area := width * height

			// part 1
			if area > sol1 {
				sol1 = area
			}

			// part 2: Check if any boundary point is inside rectangle.
			// This only works cause the polygon is "convex" enough.
			// Sorry, I was reading hints on reddit.
			xMin := min(c1.x, c2.x)
			xMax := max(c1.x, c2.x)
			yMin := min(c1.y, c2.y)
			yMax := max(c1.y, c2.y)
			var invalid bool
			for c := range boundary {
				if c.x > xMin && c.x < xMax && c.y > yMin && c.y < yMax {
					invalid = true
					break
				}
			}
			if invalid {
				continue
			}

			if area > sol2 {
				sol2 = area
			}
		}
	}

	fmt.Println("sol1:", sol1)
	fmt.Println("sol2:", sol2)
}

type Coord struct {
	x, y int
}

func NewCoord(x, y int) Coord {
	return Coord{x, y}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func CoordsBetween(c1, c2 Coord) []Coord {
	var between []Coord
	if c1.x == c2.x {
		y1, y2 := c1.y, c2.y
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for y := y1; y <= y2; y++ {
			between = append(between, NewCoord(c1.x, y))
		}
		return between
	}
	if c1.y == c2.y {
		x1, x2 := c1.x, c2.x
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		for x := x1; x <= x2; x++ {
			between = append(between, NewCoord(x, c1.y))
		}
		return between
	}
	panic("not on same row or column")
}
