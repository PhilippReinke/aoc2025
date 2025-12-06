package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
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

	fmt.Println("sol1:", part1(lines))
	fmt.Println("sol2:", part2(lines))
}

func part1(lines []string) int {
	// parse input
	var cols [][]int
	var operators []string
	for i, line := range lines {
		whitespace := regexp.MustCompile(`\s+`)
		line = strings.TrimSpace(whitespace.ReplaceAllString(line, " "))
		colsString := strings.Split(line, " ")
		if i == 0 {
			cols = make([][]int, len(colsString))
		}

		if i != len(lines)-1 {
			for colIdx, numString := range colsString {
				num, err := strconv.Atoi(numString)
				if err != nil {
					panic(err)
				}
				cols[colIdx] = append(cols[colIdx], num)
			}
		} else {
			// last row with operators
			for _, operator := range colsString {
				operators = append(operators, operator)
			}
		}
	}

	var sol int
	for i, col := range cols {
		var result int
		for numIdx, num := range col {
			switch operators[i] {
			case "+":
				result += num
			case "*":
				if numIdx == 0 {
					result = 1
				}
				result *= num
			default:
				fmt.Printf("unkown operator %q for col %v\n", operators[i], col)
				os.Exit(1)
			}
		}
		sol += result
	}

	return sol
}

func part2(lines []string) int {
	lines = rotate90Left(lines)

	var numBlocks [][]int
	var numBlock []int
	var operators []string
	for _, line := range lines {
		whitespace := regexp.MustCompile(`\s+`)
		line = strings.TrimSpace(whitespace.ReplaceAllString(line, ""))

		if line == "" {
			numBlocks = append(numBlocks, numBlock)
			numBlock = []int{}
			continue
		}

		last := string(line[len(line)-1])
		if last == "+" || last == "*" {
			operators = append(operators, last)
			line = line[:len(line)-1]
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		numBlock = append(numBlock, num)
	}
	numBlocks = append(numBlocks, numBlock)

	var sol int
	for i, numBlock := range numBlocks {
		var result int
		for numIdx, num := range numBlock {
			switch operators[i] {
			case "+":
				result += num
			case "*":
				if numIdx == 0 {
					result = 1
				}
				result *= num
			default:
				panic("unknown operation")
			}
		}
		sol += result
	}

	return sol
}

func rotate90Left(lines []string) []string {
	// Here we heavily use that all rows have equal lengths.
	maxLen := len(lines[0])

	rotated := make([]string, maxLen)
	for col := maxLen - 1; col >= 0; col-- {
		var out []byte
		for row := range len(lines) {
			out = append(out, lines[row][col])
		}
		rotated[maxLen-1-col] = string(out)
	}

	return rotated
}
