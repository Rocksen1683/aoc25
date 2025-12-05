package aoc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func ReadCommaInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("failed to open file %s: %v", path, err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ",")...)
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("failed to read file %s: %v", path, err))
	}
	return lines, nil
}

func ReadInputAsMatrix(path string) ([][]int, error) {
	lines, err := ReadInput(path)
	if err != nil {
		return nil, err
	}

	matrix := make([][]int, len(lines))
	for i, line := range lines {
		matrix[i] = make([]int, len(line))
	}
	return matrix, nil
}
