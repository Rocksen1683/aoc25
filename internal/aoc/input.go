package aoc

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("failed to open file %s: %v", path, err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("failed to read file %s: %v", path, err))
	}
	return lines, nil
}
