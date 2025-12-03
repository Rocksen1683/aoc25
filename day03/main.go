package main

import (
	"fmt"

	"github.com/Rocksen1683/aoc25/internal/aoc"
)

func main() {
	lines, err := aoc.ReadInput("day03/input.txt")
	if err != nil {
		panic(err)
	}
	pt1, pt2 := sumOfMaxJoltage(lines)
	fmt.Println("part 1:", pt1)
	fmt.Println("part 2:", pt2)
}
