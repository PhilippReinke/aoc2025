package main

import (
	"fmt"
	"io"
	"os"
	"slices"
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

	// parse grid
	var grid Grid
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	var sol1 int
	var beams []Beam
	for row := range grid {
		if row == 0 {
			beams = []Beam{NewBeam(grid.StartCol(), 1)}
			continue
		}
		if row == len(grid)-1 {
			// skip last row
			continue
		}

		var beamsUpdated []Beam
		for _, b := range beams {
			nextCols := grid.BeamDown(row, b.col)
			for _, col := range nextCols {
				beamsUpdated = append(beamsUpdated, NewBeam(col, b.count))
			}
		}
		sol1 += len(beamsUpdated) - len(beams)
		beams = dedup(beamsUpdated)
	}

	var sol2 int
	for _, b := range beams {
		sol2 += b.count
	}

	fmt.Println("sol1:", sol1)
	fmt.Println("sol2:", sol2)
}

type Beam struct {
	col   int
	count int
}

func NewBeam(col, count int) Beam {
	return Beam{
		col:   col,
		count: count,
	}
}

func dedup(in []Beam) []Beam {
	colToBeam := map[int]Beam{}
	for _, b := range in {
		otherBeam, ok := colToBeam[b.col]
		if !ok {
			colToBeam[b.col] = b
		}
		colToBeam[b.col] = NewBeam(b.col, b.count+otherBeam.count)
	}

	var out []Beam
	for _, b := range colToBeam {
		out = append(out, b)
	}

	return out
}

type Grid [][]rune

func (g Grid) Valid(r, c int) bool {
	if r < 0 || r >= len(g) {
		return false
	}
	if c < 0 || c >= len(g[r]) {
		return false
	}
	return true
}

func (g Grid) IsSplitter(row, c int) bool {
	if !g.Valid(row, c) {
		return false
	}

	return g[row][c] == '^'
}

func (g Grid) StartCol() int {
	col := slices.Index(g[0], 'S')
	if col < 0 {
		panic("start not found")
	}
	return col
}

// BeamDown return new beam columns when moving down.
func (g Grid) BeamDown(r, c int) []int {
	if g.IsSplitter(r+1, c) {
		cols := []int{}
		if g.Valid(r, c-1) {
			cols = append(cols, c-1)
		}
		if g.Valid(r, c+1) {
			cols = append(cols, c+1)
		}
		return cols
	}
	return []int{c}
}
