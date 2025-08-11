package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
		if count[num] > len(nums)/2 {
			return num
		}
	}
	return 0
}

func main() {
	fmt.Println([]int{3, 2, 3, 4, 2, 3, 1, 5, 5})
}
