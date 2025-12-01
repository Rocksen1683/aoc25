package main

import (
	"fmt"
	"strconv"
)

/*
Calculates the number of times that the dial goes to 0.
If spin is left (L) decrement, if spin is right (R) increment.
The dial starts at 50.
If the dial goes below 0, it wraps around to 99.
*/
func Part1(input []string) int {
	numZeroes := 0
	dial := 50

	for _, line := range input {
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(fmt.Sprintf("failed to convert %s to int: %v", line[1:], err))
		}
		num %= 100
		dir := line[0]
		switch dir {
		case 'L':
			dial = (dial - num) % 100
			if dial < 0 {
				dial += 100
			}
		case 'R':
			dial = (dial + num) % 100
		default:
			panic(fmt.Sprintf("invalid direction %q in line %q", dir, line))
		}

		if dial == 0 {
			numZeroes++
		}

	}

	return numZeroes
}

/*
Same as part 1, but we also want to count the number of times that the dial goes over 0 and not necessarily stop at 0. 
*/
func Part2(input []string) int {
	numZeroes := 0
	dial := 50 

	for _, line := range input {
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(fmt.Errorf("failed to convert %s to int: %v", line[1:], err));
		}

		dir := line[0]
		var step int
		switch dir {
		case 'L':
			step = -1
		case 'R':
			step = 1
		default:
			panic(fmt.Sprintf("invalid direction %q in line %q", dir, line))
		}

		for i := 0; i < num; i++ {
			dial += step

			if dial < 0 {
				dial = 99
			} else if dial > 99 {
				dial = 0
			}

			if dial == 0 {
				numZeroes++
			}
		}
	}

	return numZeroes
}