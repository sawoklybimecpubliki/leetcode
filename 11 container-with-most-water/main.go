package main

import "fmt"

func maxArea(height []int) int {

	left := 0
	right := len(height) - 1
	maxArea := 0

	for left < right {
		width := right - left
		h := height[left]
		if height[right] < h {
			h = height[right]
		}
		area := h * width
		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
	/*	var maxHeight, square int

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
		return square*/
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
