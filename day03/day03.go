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

	var sol1 int
	for _, line := range lines {
		// find max before last element
		maxIdx := 0
		for i := 1; i < len(line)-1; i++ {
			if line[maxIdx] < line[i] {
				maxIdx = i
			}
		}
		// find max after previous max
		maxIdx2 := maxIdx + 1
		for i := maxIdx + 2; i < len(line); i++ {
			if line[maxIdx2] < line[i] {
				maxIdx2 = i
			}
		}

		first := string(line[min(maxIdx, maxIdx2)])
		second := string(line[max(maxIdx, maxIdx2)])
		joltage, err := strconv.Atoi(first + second)
		if err != nil {
			panic(err)
		}

		sol1 += joltage
	}

	// part 2
	var sol2 int
	for _, line := range lines {
		joltageString := part2(line, 12)
		joltage, err := strconv.Atoi(joltageString)
		if err != nil {
			panic(err)
		}
		sol2 += joltage
	}

	fmt.Println("sol1:", sol1)
	fmt.Println("sol2:", sol2)
}

func part2(line string, k int) string {
	var result []byte
	startIdx := 0

	for len(result) < k {
		remaining := k - len(result)
		end := len(line) - remaining

		// find max between previous max and "moving" end
		maxIdx := startIdx
		for i := startIdx + 1; i <= end; i++ {
			if line[i] > line[maxIdx] {
				maxIdx = i
			}
		}

		result = append(result, line[maxIdx])
		startIdx = maxIdx + 1
	}

	return string(result)
}
