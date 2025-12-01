package main

import (
	"fmt"

	"github.com/Rocksen1683/aoc25/internal/aoc"
)

func main() {
	lines, err := aoc.ReadInput("day01/input.txt")
	fmt.Println(lines)
	if err != nil {
		panic(err)
	}
	fmt.Println(Part1(lines))
}
