package main

import (
	"fmt"
	"io"
	"os"
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

	// read grid
	var grid Grid
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	// part 1
	sol1 := grid.RemovableRolls()

	// part 2
	var sol2 int
	cnt := sol1
	for cnt > 0 {
		// remove removable rolls
		for row := range grid {
			for col := range grid[row] {
				if grid.IsRoll(row, col) && grid.AdjacentRolls(row, col) < 4 {
					grid.Remove(row, col)
					sol2++
				}
			}
		}

		// compute new count of removable rolls
		cnt = grid.RemovableRolls()
	}

	fmt.Println("sol1:", sol1)
	fmt.Println("sol2:", sol2)
}

type Grid [][]rune

func (g Grid) IsRoll(r, c int) bool {
	if r < 0 || r >= len(g) {
		return false
	}
	if c < 0 || c >= len(g[r]) {
		return false
	}
	return g[r][c] == '@'
}

func (g Grid) AdjacentRolls(r, c int) int {
	var cnt int
	if g.IsRoll(r, c+1) {
		cnt++
	}
	if g.IsRoll(r, c-1) {
		cnt++
	}
	if g.IsRoll(r+1, c+1) {
		cnt++
	}
	if g.IsRoll(r+1, c-1) {
		cnt++
	}
	if g.IsRoll(r-1, c+1) {
		cnt++
	}
	if g.IsRoll(r-1, c-1) {
		cnt++
	}
	if g.IsRoll(r+1, c) {
		cnt++
	}
	if g.IsRoll(r-1, c) {
		cnt++
	}
	return cnt
}

func (g Grid) RemovableRolls() int {
	var cnt int
	for row := range g {
		for col := range g[row] {
			if g.IsRoll(row, col) && g.AdjacentRolls(row, col) < 4 {
				cnt++
			}
		}
	}
	return cnt
}

func (g Grid) Remove(r, c int) {
	g[r][c] = '.'
}
