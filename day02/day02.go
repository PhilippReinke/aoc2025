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
	line := strings.Split(string(dataTrimmed), "\n")[0]
	idRanges := strings.Split(line, ",")

	var sol1, sol2 int
	for _, idRange := range idRanges {
		numStrings := strings.Split(idRange, "-")
		if len(numStrings) != 2 {
			panic("invalid range")
		}
		num1String := numStrings[0]
		num2String := numStrings[1]
		num1, err := strconv.Atoi(num1String)
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(num2String)
		if err != nil {
			panic(err)
		}

		for num := num1; num <= num2; num++ {
			if !validPart1(num) {
				sol1 += num
			}
			if !validPart2(num) {
				sol2 += num
			}
		}
	}

	fmt.Println("sol1:", sol1)
	fmt.Println("sol2:", sol2)
}

func validPart1(num int) bool {
	numString := strconv.Itoa(num)

	if len(numString) == 0 {
		panic("zero length")
	}
	if len(numString)%2 == 1 {
		// uneven is always valid
		return true
	}
	half := len(numString) / 2

	return numString[:half] != numString[half:]
}

func validPart2(num int) bool {
	numString := strconv.Itoa(num)
	numLen := len(numString)

	// take first i characters and check if repeating them produces original
	// number
	for i := 1; i <= numLen/2; i++ {
		if numLen%i != 0 {
			// number length not a multiple of attempted split length
			continue
		}
		split := numString[:i]

		repeated := strings.Repeat(split, numLen/i)
		if repeated == numString {
			return false
		}
	}

	return true
}
