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
	blocks := strings.Split(string(dataTrimmed), "\n\n")

	var sol1, sol2 int
	for _, block := range blocks {
		for _, line := range strings.Split(block, "\n") {
			if line[len(line)-1] == ':' {
				// TODO
			}
			// TODO
		}
	}

	fmt.Println("sol1:", sol1)
	fmt.Println("sol2:", sol2)
}

var Shape [][]bool
