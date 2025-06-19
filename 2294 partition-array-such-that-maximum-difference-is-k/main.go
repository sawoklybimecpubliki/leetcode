package main

import (
	"fmt"
	"sort"
)

func partitionArray(nums []int, k int) int {
	var res int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[j]-nums[i] > k {
				i = j - 1
				break
			}
			if j == len(nums)-1 {
				return res + 1
			}
		}
		res++
	}
	return res

	/*	func partitionArray(nums []int, k int) int {
			res := 1
			left := 0
			sort.Ints(nums)
			for i := 1; i < len(nums); i++ {
			if nums[i]-nums[left] > k {
			res++
			left = i
		}
		}
			return res
		}*/
}

func main() {
	fmt.Println(partitionArray([]int{3, 6, 1, 2, 5}, 2))
}
