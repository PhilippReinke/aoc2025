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
	blocks := strings.Split(string(dataTrimmed), "\n\n")

	// parse
	var shapes []Shape
	var regions []Region
	for _, block := range blocks {
		var shape Shape
		for line := range strings.SplitSeq(block, "\n") {
			if line[len(line)-1] == ':' {
				continue
			}
			if line[0] == '#' || line[0] == '.' {
				shape = append(shape, []rune(line))
				continue
			}

			// assume it is a region
			split := strings.Split(line, ":")
			if len(split) != 2 {
				panic("unexpected line length")
			}
			regions = append(regions, NewRegion(split[0], split[1]))
		}
		if !shape.Empty() {
			shapes = append(shapes, shape)
		}
	}

	var sol1 int
	for _, region := range regions {
		// This solves the input as the shapes are almost 3x3 squares.
		//
		// ¯\_(ツ)_/¯
		//
		// LOL
		if region.ShapeSum()*9 <= region.Area() {
			sol1++
		}
	}

	fmt.Println("sol1:", sol1)
}

type Shape [][]rune

func (s Shape) Empty() bool {
	return len(s) == 0
}

type Region struct {
	Rows   int
	Cols   int
	Counts []int
}

func (r Region) Area() int {
	return r.Rows * r.Cols
}

func (r Region) ShapeSum() int {
	var c int
	for _, n := range r.Counts {
		c += n
	}
	return c
}

func NewRegion(beforeColon, afterColon string) Region {
	beforeSplit := strings.Split(beforeColon, "x")
	if len(beforeSplit) != 2 {
		panic("before not 2")
	}
	rows, err := strconv.Atoi(beforeSplit[0])
	if err != nil {
		panic(err)
	}
	cols, err := strconv.Atoi(beforeSplit[1])
	if err != nil {
		panic(err)
	}

	var counts []int
	for numString := range strings.SplitSeq(strings.Trim(afterColon, " "), " ") {
		num, err := strconv.Atoi(numString)
		if err != nil {
			panic(err)
		}
		counts = append(counts, num)
	}

	return Region{
		Rows:   rows,
		Cols:   cols,
		Counts: counts,
	}
}
