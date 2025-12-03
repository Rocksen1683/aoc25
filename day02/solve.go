package main

import (
	"strconv"
	"strings"
)

func isInvalidNumber(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	half := len(s) / 2
	return strings.Contains(s[half:], s[:half])
}

func splitStringIntoChunks(s string, n int) []string {
	chunks := make([]string, 0, len(s)/n)
	for i := 0; i < len(s); i += n {
		chunks = append(chunks, s[i:i+n])
	}
	return chunks
}

func checkInvalidChunks(chunks []string) bool {
	start := chunks[0]
	for _, chunk := range chunks[1:] {
		if chunk != start {
			return true
		}
	}
	return false
}

func isInvalidSubstring(s string) bool {
	length := len(s)
	for i := 1; i < length; i++ {
		if length%i != 0 {
			continue
		}
		chunks := splitStringIntoChunks(s, i)
		if !checkInvalidChunks(chunks) {
			return true
		}
	}
	return false
}

func findInvalidRange(input []string, part2 bool) int {
	invalidNumbers := 0

	for _, line := range input {
		parts := strings.Split(line, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		for i := start; i <= end; i++ {
			if part2 {
				if isInvalidSubstring(strconv.Itoa(i)) {
					invalidNumbers += i
				}
			} else if !part2 && isInvalidNumber(strconv.Itoa(i)) {
				invalidNumbers += i
			}
		}
	}

	return invalidNumbers
}
