package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	var result []int
	var xCount, yCount int
	dir := 1
	xMax := len(matrix[0])
	yMax := len(matrix)
	lenMatrix := xMax * yMax
	xMax--
	yMax--
	for len(result) < lenMatrix {
		switch dir {
		case 1:
			for i := xCount; i <= xMax-xCount; i++ {
				result = append(result, matrix[yCount][i])
			}
			dir++
		case 2:
			for j := yCount + 1; j <= yMax-yCount; j++ {
				result = append(result, matrix[j][xMax-xCount])
			}
			dir++
		case 3:
			for i := xMax - xCount - 1; i >= xCount; i-- {
				result = append(result, matrix[yMax-yCount][i])
			}
			dir++
		case 4:
			for j := yMax - yCount - 1; j >= yCount+1; j-- {
				result = append(result, matrix[j][xCount])
			}
			dir++
		case 5:
			dir = 1
			xCount++
			yCount++
		}
	}
	// 0 - max -> max - 0 -> 0 max-1 -> max-1 - 1
	return result
}

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
