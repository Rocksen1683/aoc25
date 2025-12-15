package main

import "sort"

func findNumFreshIDs(intervals [][2]int, freshIDs []int) int {
	numFreshIDs := 0

	for _, freshID := range freshIDs {
		for _, interval := range intervals {
			if freshID >= interval[0] && freshID <= interval[1] {
				numFreshIDs++
				break
			}
		}
	}
	return numFreshIDs
}

func findNumFreshIntervals(intervals [][2]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	total := 0
	curStart, curEnd := intervals[0][0], intervals[0][1]

	for k := 1; k < len(intervals); k++ {
		start, end := intervals[k][0], intervals[k][1]

		if start > curEnd+1 {
			total += curEnd - curStart + 1
			curStart, curEnd = start, end
			continue
		}

		if end > curEnd {
			curEnd = end
		}
	}

	total += curEnd - curStart + 1
	return total
}