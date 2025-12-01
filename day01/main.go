package main

import (
	"fmt"

	"github.com/Rocksen1683/aoc25/internal/aoc"
)

func main() {
	lines, err := aoc.ReadInput("day01/input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("part 1:", Part1(lines))
	fmt.Println("part 2:", Part2(lines))
}
