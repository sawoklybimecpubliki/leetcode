package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	for row := len(triangle) - 2; row >= 0; row-- {
		for col := 0; col <= row; col++ {
			belowLeft := triangle[row+1][col]
			belowRight := triangle[row+1][col+1]
			triangle[row][col] += min(belowLeft, belowRight)
		}
	}
	return triangle[0][0]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}
