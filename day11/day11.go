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

	// parse data
	m := map[string][]string{}
	for _, line := range lines {
		splitColon := strings.Split(line, ":")
		if len(splitColon) != 2 {
			panic("split colon not 2")
		}
		key := splitColon[0]
		rest := strings.Trim(splitColon[1], " ")
		m[key] = strings.Split(rest, " ")
	}

	// part 1
	sol1 := countPaths("you", m)

	// part 2
	cache := map[string]int{}
	sol2 := countPathsPart2("svr", m, false, false, cache)

	fmt.Println("sol1:", sol1)
	fmt.Println("sol2:", sol2)
}

func countPaths(node string, connections map[string][]string) int {
	if node == "out" {
		return 1
	}

	count := 0
	if nextNodes, ok := connections[node]; ok {
		for _, nextNode := range nextNodes {
			count += countPaths(nextNode, connections)
		}
	}

	return count
}

func countPathsPart2(node string, connections map[string][]string, visitedDac, visitedFft bool, cache map[string]int) int {
	if node == "dac" {
		visitedDac = true
	}
	if node == "fft" {
		visitedFft = true
	}

	if node == "out" {
		if visitedDac && visitedFft {
			return 1
		}
		return 0
	}

	key := fmt.Sprintf("%s-%t-%t", node, visitedDac, visitedFft)
	if val, ok := cache[key]; ok {
		return val
	}

	count := 0
	if nextNodes, ok := connections[node]; ok {
		for _, next := range nextNodes {
			count += countPathsPart2(next, connections, visitedDac, visitedFft, cache)
		}
	}

	cache[key] = count
	return count
}
