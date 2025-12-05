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
	fmt.Println("rows:", len(matrix))
	fmt.Println("cols:", len(matrix[0]))
}
