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

func findMaxIdx(lines []rune, start int, end int) int {
	maxNum := 0
	maxIdx := 0

	for i := start; i <= end; i++ {
		convNum, err := strconv.Atoi(string(lines[i]))
		if err != nil {
			panic(err)
		}
		if convNum > maxNum {
			maxNum = convNum
			maxIdx = i
		}
		if maxNum == 9 {
			break
		}
	}

	return maxIdx
}

func findMaxJoltageTwelve(lines string) int {
	length := len(lines)
	var chars []rune
	st := 0

	for len(chars) < 12 {
		maxIdx := findMaxIdx([]rune(lines), st, length-12+len(chars))
		maxNum := lines[maxIdx]
		chars = append(chars, rune(maxNum))
		st = maxIdx + 1
	}

	convNum, err := strconv.Atoi(string(chars))
	if err != nil {
		panic(err)
	}
	return convNum
}

func sumOfMaxJoltage(lines []string) (int, int) {
	//part 1
	partOneSum := 0
	for _, line := range lines {
		partOneSum += findMaxJoltage(line)
	}

	//part 2
	partTwoSum := 0
	for _, line := range lines {
		partTwoSum += findMaxJoltageTwelve(line)
	}

	return partOneSum, partTwoSum
}
