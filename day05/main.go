package main

import (
	"fmt"

	"github.com/Rocksen1683/aoc25/internal/aoc"
)

func main() {
	intervals, freshIDs, err := aoc.ReadInputAsIntervals("day05/input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("part 1:", findNumFreshIDs(intervals, freshIDs))
	fmt.Println("part 2:", findNumFreshIntervals(intervals))
}
