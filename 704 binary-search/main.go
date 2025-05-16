package main

import "fmt"

func search(nums []int, target int) int {
	for i, n := range nums {
		if n == target {
			return i
		}
		if n > target {
			return -1
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}
