package main

import (
	"fmt"
	"io"
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
	var blankLine bool
	var ranges []Range
	var ingredientIDs []int
	for _, line := range lines {
		if line == "" {
			blankLine = true
			continue
		}

		if !blankLine {
			// range
			split := strings.Split(line, "-")
			if len(split) != 2 {
				panic("range length invalid")
			}
			begin, err := strconv.Atoi(split[0])
			if err != nil {
				panic(err)
			}
			end, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}
			ranges = append(ranges, Range{begin, end})
		} else {
			// ingredient ID
			id, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			ingredientIDs = append(ingredientIDs, id)
		}
	}

	// part 1
	var sol1 int
	for _, id := range ingredientIDs {
		for _, ran := range ranges {
			if id >= ran.Start && id <= ran.End {
				sol1++
				break
			}
		}
	}

	// part 2
	var sol2 int
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})
	for i := 0; i < len(ranges); i++ {
		for j := i + 1; j < len(ranges); j++ {
			if ranges[j].Start <= ranges[i].End {
				ranges[j] = Range{
					Start: ranges[i].End + 1,
					End:   ranges[j].End,
				}
			}
		}
	}
	for _, ran := range ranges {
		length := ran.End - ran.Start + 1
		if length > 0 {
			sol2 += length
		}
	}

	fmt.Println("sol1:", sol1)
	fmt.Println("sol2:", sol2)
}

type Range struct {
	Start, End int
}
