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

	// Locate every operator column from right-to-left; each operator marks the
	// start of a problem whose columns extend to the next operator (or the
	// right edge).
	var opCols []int
	for col := width - 1; col >= 0; col-- {
		if ch := opLine[col]; ch == '+' || ch == '*' {
			opCols = append(opCols, col)
		}
	}
	if len(opCols) == 0 {
		return nil, fmt.Errorf("no operations found in %s", path)
	}

	cephalopods := make([]Cephalopod, 0, len(opCols))

	rightBound := width
	for _, opCol := range opCols {
		cp := Cephalopod{
			Operation: string(opLine[opCol]),
		}

		for col := opCol; col < rightBound; col++ {
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
				return nil, fmt.Errorf("failed to parse number at column %d: %v", col, err)
			}
			cp.Numbers = append(cp.Numbers, num)
		}

		if len(cp.Numbers) == 0 {
			return nil, fmt.Errorf("no numbers found for operator at column %d", opCol)
		}

		cephalopods = append(cephalopods, cp)
		rightBound = opCol
	}

	return cephalopods, nil
}
