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

	pos := 50
	var zeroCount, zeroCount2 int
	for _, line := range lines {
		numString := line[1:]
		num, err := strconv.Atoi(numString)
		if err != nil {
			panic("cannot parse number")
		}

		var zeroPass int
		direction := line[0]
		switch direction {
		case 'L':
			pos, zeroPass = calcPos(pos, -num)
		case 'R':
			pos, zeroPass = calcPos(pos, num)
		default:
			panic("unknown direction")
		}

		if pos == 0 {
			zeroCount++
		}
		zeroCount2 += zeroPass
	}

	fmt.Println("sol1:", zeroCount)
	fmt.Println("sol2:", zeroCount2)
}

func calcPos(cur, rot int) (int, int) {
	if rot == 0 {
		return cur, 0
	}
	newPos := ((cur+rot)%100 + 100) % 100

	fullCycles := abs(rot) / 100
	zeroPass := fullCycles

	// There might be a partial passings, e.g. from 99 going R102 has two zero
	// passes.
	partial := abs(rot) % 100
	if rot > 0 {
		stepsToZero := (100 - cur) % 100
		if stepsToZero == 0 {
			stepsToZero = 100
		}
		if partial >= stepsToZero {
			zeroPass++
		}
	} else if rot < 0 {
		stepsToZero := cur % 100
		if stepsToZero == 0 {
			stepsToZero = 100
		}
		if partial >= stepsToZero {
			zeroPass++
		}
	}

	return newPos, zeroPass
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
