package aoc

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func ReadInputAsMatrix(path string) ([][]string, error) {
	lines, err := ReadInput(path)
	if err != nil {
		return nil, err
	}

	matrix := make([][]string, len(lines))
	for i, line := range lines {
		matrix[i] = strings.Split(line, "")
	}
	return matrix, nil
}

func ReadInputAsIntervals(path string) ([][2]int, []int, error) {
	lines, err := ReadInput(path)
	if err != nil {
		return nil, nil, err
	}

	intervals := make([][2]int, 0)
	freshIDs := make([]int, 0)

	for _, line := range lines {
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			start, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, nil, err
			}
			end, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, nil, err
			}
			intervals = append(intervals, [2]int{start, end})
		} else if line == "" {
			continue
		} else {
			id, err := strconv.Atoi(line)
			if err != nil {
				return nil, nil, err
			}
			freshIDs = append(freshIDs, id)
		}
	}
	return intervals, freshIDs, nil
}

func ReadInputAsCephalopod(path string) ([]Cephalopod, error) {
	lines, err := ReadInput(path)
	if err != nil {
		return nil, err
	}

	if len(lines) == 0 {
		return nil, fmt.Errorf("no lines in input")
	}

	// Use Fields to handle multiple spaces and split properly
	firstLineParts := strings.Fields(lines[0])
	cephalopods := make([]Cephalopod, len(firstLineParts))

	for lineIdx, line := range lines {
		parts := strings.Fields(line)
		for idx := 0; idx < len(parts); idx++ {
			if lineIdx == len(lines)-1 {
				cephalopods[idx].Operation = parts[idx]
			} else {
				numbers, err := strconv.Atoi(parts[idx])
				if err != nil {
					return nil, err
				}
				cephalopods[idx].Numbers = append(cephalopods[idx].Numbers, numbers)
			}
		}
	}
	return cephalopods, nil
}

func ReadInputAsCephalopodRightToLeft(path string) ([]Cephalopod, error) {
	lines, err := ReadInput(path)
	if err != nil {
		return nil, err
	}

	if len(lines) == 0 {
		return nil, fmt.Errorf("no lines in input")
	}

	numDataLines := len(lines) - 1
	if numDataLines == 0 {
		return nil, fmt.Errorf("no data lines")
	}

	width := 0
	for _, line := range lines {
		if len(line) > width {
			width = len(line)
		}
	}

	// Pad all lines to equal width.
	padded := make([]string, len(lines))
	for i, line := range lines {
		if len(line) < width {
			line = line + strings.Repeat(" ", width-len(line))
		}
		padded[i] = line
	}

	dataLines := padded[:numDataLines]
	opLine := padded[len(padded)-1]

	type block struct {
		cols []int 
	}
	var blocks []block

	inBlock := false
	current := block{}

	for col := width - 1; col >= 0; col-- {
		allSpace := opLine[col] == ' '
		if allSpace {
			for _, row := range dataLines {
				if row[col] != ' ' {
					allSpace = false
					break
				}
			}
		}

		if allSpace {
			if inBlock {
				blocks = append(blocks, current)
				current = block{}
				inBlock = false
			}
			continue
		}

		if !inBlock {
			inBlock = true
		}
		current.cols = append(current.cols, col)
	}
	if inBlock {
		blocks = append(blocks, current)
	}

	cephalopods := make([]Cephalopod, 0, len(blocks))

	for _, blk := range blocks {
		var cp Cephalopod

		for _, col := range blk.cols {
			if opLine[col] != ' ' {
				cp.Operation = string(opLine[col])
				break
			}
		}
		if cp.Operation == "" {
			return nil, fmt.Errorf("missing operation for block with columns %v", blk.cols)
		}

		for _, col := range blk.cols {
			var digits []byte
			for _, row := range dataLines {
				ch := row[col]
				if ch >= '0' && ch <= '9' {
					digits = append(digits, ch)
				}
			}
			if len(digits) == 0 {
				continue
			}
			num, err := strconv.Atoi(string(digits))
			if err != nil {
				return nil, fmt.Errorf("failed to parse number in column %d: %v", col, err)
			}
			cp.Numbers = append(cp.Numbers, num)
		}

		cephalopods = append(cephalopods, cp)
	}

	return cephalopods, nil
}
