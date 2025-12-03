package main

import (
	"strconv"
)

func findMaxJoltage(lines string) int {
	length := len(lines)
	maxNum := 0
	maxIdx := 0

	for i := 0; i < length-1; i++ {
		convNum, err := strconv.Atoi(string(lines[i]))
		if err != nil {
			panic(err)
		}
		if convNum > maxNum {
			maxNum = convNum
			maxIdx = i
		}

		//break if we hit 9
		if maxNum == 9 {
			break
		}
	}

	minNum := 0

	for i := length - 1; i > maxIdx; i-- {
		convNum, err := strconv.Atoi(string(lines[i]))
		if err != nil {
			panic(err)
		}
		if convNum > minNum && i != maxIdx {
			minNum = convNum
		}

		if minNum == 9 {
			break
		}
	}

	return maxNum*10 + minNum
}

func sumOfMaxJoltage(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += findMaxJoltage(line)
	}

	return sum
}
