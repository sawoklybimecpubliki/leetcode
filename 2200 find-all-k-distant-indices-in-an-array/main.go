package main

import "fmt"

func findKDistantIndices(nums []int, key int, k int) []int {
	var out []int
	var index []int
	for i, n := range nums {
		if n == key {
			index = append(index, i)
		}
	}
	for i := range nums {
		for _, j := range index {
			if abs(i, j) <= k {
				out = append(out, i)
				break
			}
		}
	}
	return out
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	fmt.Println(findKDistantIndices([]int{2, 2, 2, 2, 2}, 2, 2))
}
