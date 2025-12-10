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

	var sol1, sol2 int
	for _, line := range lines {
		wantState, buttons, _ := parse(line)

		fmt.Println(wantState)
		fmt.Println(buttons)
	}

	fmt.Println("sol1:", sol1)
	fmt.Println("sol2:", sol2)
}

type State []bool
type Button []int

func (s State) Copy() State {
	cpy := make(State, len(s))
	for i := range s {
		cpy[i] = s[i]
	}
	return cpy
}

func (s State) Apply(b Button) State {
	newState := s.Copy()
	for _, i := range b {
		newState[i] = !newState[i]
	}
	return newState
}

func (s State) Equal(other State) bool {
	if len(s) != len(other) {
		return false
	}
	for i := range s {
		if s[i] != other[i] {
			return false
		}
	}
	return true
}

// parse parses line to desired states, buttons, joltage
func parse(line string) (State, []Button, []int) {
	var states State
	var buttons []Button

	for seg := range strings.SplitSeq(line, " ") {
		switch seg[0] {
		case '[':
			for _, r := range seg[1 : len(seg)-1] {
				on := r == '#'
				states = append(states, on)
			}
		case '(':
			var button []int
			for numString := range strings.SplitSeq(seg[1:len(seg)-1], ",") {
				num, err := strconv.Atoi(numString)
				if err != nil {
					panic(err)
				}
				button = append(button, num)
			}
			buttons = append(buttons, button)
		case '{':
			continue
		default:
			fmt.Println("unknown seg start:", seg)
			panic("failed to parse line")
		}
	}

	return states, buttons, []int{}
}
