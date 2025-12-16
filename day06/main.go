package main

import (
	"fmt"

	"github.com/Rocksen1683/aoc25/internal/aoc"
)

func main() {
	cephalopods, err := aoc.ReadInputAsCephalopod("day06/input.txt")
	if err != nil {
		panic(err)
	}

	rightToLeftCephalopods, err := aoc.ReadInputAsCephalopodRightToLeft("day06/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("part 1:", calculateCephalopod(cephalopods))
	fmt.Println("part 2:", calculateCephalopod(rightToLeftCephalopods))
}
