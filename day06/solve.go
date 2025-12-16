package main

import "github.com/Rocksen1683/aoc25/internal/aoc"

func calculateCephalopodValue(cephalopod aoc.Cephalopod) int {
	value := 0

	for i, number := range cephalopod.Numbers {
		if i == 0 {
			value = number
			continue
		}

		switch cephalopod.Operation {
		case "+":
			value += number
		case "*":
			value *= number
		}
	}

	return value
}

func calculateCephalopod(cephalopods []aoc.Cephalopod) int {
	totalSum := 0
	for _, cephalopod := range cephalopods {
		totalSum += calculateCephalopodValue(cephalopod)
	}
	return totalSum
}
