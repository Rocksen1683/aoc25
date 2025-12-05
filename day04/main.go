package main

import (
	"fmt"

	"github.com/Rocksen1683/aoc25/internal/aoc"
)

func main() {
	matrix, err := aoc.ReadInputAsMatrix("day04/input.txt")
	if err != nil {
		panic(err)
	}
	copyMatrix := make([][]string, len(matrix))
	for i := range matrix {
		copyMatrix[i] = make([]string, len(matrix[i]))
		copy(copyMatrix[i], matrix[i])
	}

	fmt.Println("rows:", len(matrix))
	fmt.Println("cols:", len(matrix[0]))
	fmt.Println("part 1:", numAdjPaperRolls(copyMatrix))
	fmt.Println("part 2:", numPaperRollsToRemove(matrix))
}
