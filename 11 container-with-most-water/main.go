package main

import "fmt"

func maxArea(height []int) int {
	var maxHeight, square int

	for i := len(height) - 1; i > 0; i-- {

		if maxHeight < height[i] {
			maxHeight = height[i]
			for j := 0; j < i; j++ {
				if height[i] > height[j] {

					if (i-j)*height[j] > square {
						square = (i - j) * height[j]
					}

				} else {

					if (i-j)*height[i] > square {
						square = (i - j) * height[i]
					}

				}

			}
		}

	}
	return square
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
