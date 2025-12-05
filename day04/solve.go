package main

func numAdjPaperRolls(matrix [][]string) int {
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	numPaperRolls := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != "@" {
				continue
			}
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
			if numAdj < 4 {
				numPaperRolls++
			}
		}
	}
	return numPaperRolls
}
