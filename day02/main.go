package main

import (
	"fmt"

	"github.com/Rocksen1683/aoc25/internal/aoc"
)

func main() {
	lines, err := aoc.ReadCommaInput("day02/input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("part 1:", findInvalidRange(lines, false))
	fmt.Println("part 2:", findInvalidRange(lines, true))
}
