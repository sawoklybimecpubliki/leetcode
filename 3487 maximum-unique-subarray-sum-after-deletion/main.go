package main

import (
	"fmt"
	"math"
)

func maxSum(nums []int) int {
	mn := math.MinInt32
	pos := false
	for _, val := range nums {
		if val > 0 {
			pos = true
			break
		}
		if val > mn {
			mn = val
		}
	}

	if !pos {
		return mn
	}

	seen := make(map[int]bool)
	sum := 0

	for _, val := range nums {
		if val >= 0 {
			if !seen[val] {
				if val >= 0 {
					sum += val
				}
			}
			seen[val] = true
		}
	}

	return sum
}

func main() {
	fmt.Println(maxSum([]int{1, 2, -1, -2, 1, 0, -1}))
}
