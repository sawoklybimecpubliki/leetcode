package main

import "fmt"

func setZeroes(matrix [][]int) {
	m, n := len(matrix[0]), len(matrix)
	var arr [][]int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				arr = append(arr, []int{i, j})
			}
		}
	}
	fmt.Println(matrix)
	for _, num := range arr {
		for i := 0; i < n; i++ {
			matrix[i][num[1]] = 0
		}
		for j := 0; j < m; j++ {
			matrix[num[0]][j] = 0
		}
	}
	fmt.Println(matrix)
}

func main() {
	setZeroes([][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}})
}
