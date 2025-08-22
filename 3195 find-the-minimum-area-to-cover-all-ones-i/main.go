package main

import "fmt"

func minimumArea(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	minRow, minCol := rows, cols
	maxRow, maxCol := -1, -1

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				if i < minRow {
					minRow = i
				}
				if i > maxRow {
					maxRow = i
				}
				if j < minCol {
					minCol = j
				}
				if j > maxCol {
					maxCol = j
				}
			}
		}
	}
	if maxRow == -1 {
		return 0
	}
	return (maxRow - minRow + 1) * (maxCol - minCol + 1)
}

func main() {
	fmt.Println(minimumArea([][]int{{0, 1, 0}, {1, 0, 1}}))
}
