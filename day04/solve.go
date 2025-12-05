package main

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func checkNumAdj(matrix [][]string, i, j int) int {
	numAdj := 0
	for _, dir := range dirs {
		ni, nj := i+dir[0], j+dir[1]
		if ni < 0 || ni >= len(matrix) || nj < 0 || nj >= len(matrix[i]) {
			continue
		}
		if matrix[ni][nj] == "@" {
			numAdj++
		}
	}
	return numAdj
}

func numAdjPaperRolls(matrix [][]string) int {
	numPaperRolls := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != "@" {
				continue
			}

			if checkNumAdj(matrix, i, j) < 4 {
				numPaperRolls++
			}
		}
	}
	return numPaperRolls
}

func numPaperRollsToRemove(matrix [][]string) int {
	numPaperRolls := 0
	for {
		var toRemove [][2]int //indexes of paper rolls to remove

		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[i]); j++ {
				if matrix[i][j] != "@" {
					continue
				}
				if checkNumAdj(matrix, i, j) < 4 {
					toRemove = append(toRemove, [2]int{i, j})
				}
			}
		}

		if len(toRemove) == 0 {
			break
		}

		numPaperRolls += len(toRemove)
		for _, pos := range toRemove {
			matrix[pos[0]][pos[1]] = "."
		}

	}

	return numPaperRolls
}
