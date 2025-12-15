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

	for _, cephalopod := range cephalopods {
		fmt.Println(cephalopod.Numbers)
		fmt.Println(cephalopod.Operation)
	}

}
