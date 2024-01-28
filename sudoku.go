package main

import (
	"fmt"
	"slices"
)

type position struct {
	x int
	y int
}

func SolveSudoku(puzzle [][]int) [][]int {
	// create puzzle to be modified for backtracking
	solved := make([][]int, len(puzzle))
	for i := range puzzle {
		solved[i] = make([]int, len(puzzle[i]))
		copy(solved[i], puzzle[i])
	}

	// initialize map of what has already been tried
	trials := map[position]int{}

	// initialize list of what positions were tried
	hitlist := []position{}

	// go through puzzle to find and attempt to solve empty cells
	i := 0
	for i < len(puzzle) {
		j := 0
		for j < len(puzzle[i]) {
			// only need to solve if cell is unsolved
			if puzzle[i][j] == 0 {
				// check what cell could be
				old, ok := trials[position{i, j}]
				var nextCandidate int
				if ok {
					nextCandidate = old + 1
				} else {
					nextCandidate = 1
				}

				used := []int{}

				// check which numbers are candidates in the col
				for y := 0; y < 9; y++ {
					if solved[y][j] != 0 {
						used = append(used, solved[y][j])
					}
				}

				// check which numbers are candidates in the row
				for x := 0; x < 9; x++ {
					if solved[i][x] != 0 {
						used = append(used, solved[i][x])
					}
				}

				// check which numbers are candidates in the block
				for x := i - (i % 3); x <= i+2-(i%3) && x < len(puzzle); x++ {
					for y := j - (j % 3); y <= j+2-(j%3) && y < len(puzzle[i]); y++ {
						used = append(used, solved[x][y])
					}
				}

				for slices.Contains(used, nextCandidate) {
					nextCandidate++
				}

				// check if backtracking is required
				if nextCandidate >= 10 {
					// implement backtracking
					solved[i][j] = 0
					trials[position{i, j}] = 0

					i = hitlist[len(hitlist)-1].x
					j = hitlist[len(hitlist)-1].y
					if len(hitlist) > 0 {
						hitlist = hitlist[:len(hitlist)-1]
					}
				} else {
					solved[i][j] = nextCandidate
					trials[position{i, j}] = nextCandidate
					hitlist = append(hitlist, position{i, j})
					j++
				}
			} else {
				j++
			}
		}
		i++
	}

	fmt.Println(hitlist)
	fmt.Println(trials)

	return solved
}
