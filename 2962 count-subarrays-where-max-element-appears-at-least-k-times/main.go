package main

import (
	"fmt"
)

func countSubarrays(nums []int, k int) int64 {
	m := 0
	leftIndex := 0
	var index []int
	for _, num := range nums {
		if m < num {
			m = num
		}
	}
	for i, num := range nums {
		if m == num {
			index = append(index, i)
		}
	}
	result := 0
	fmt.Println(index)
	for i := 0; i+k-1 < len(index); i++ {
		fmt.Println("left index: ", index[i], "right index: ", index[i+k-1])
		result = result + (len(nums)-index[i+k-1])*(index[i]-leftIndex+1)
		leftIndex = index[i] + 1
	}
	/*	j := 0
		result := 0
		count := 0
		for i := 0; i < len(nums); {
			if nums[j] == m {
				if count == 0 {
					leftIndex = j
					fmt.Println("left: ", leftIndex)
				}
				count++
			}
			if count == k {
				result = result + (len(nums)-j)*(leftIndex-i+1)
				count = 0
				i = leftIndex + 1
				j = i
			}
			if j != len(nums)-1 {
				j++
			} else {
				i = i + 1
				j = i
				count = 0
			}

		}*/
	return int64(result)
}

func main() {
	fmt.Println("answer: ", countSubarrays([]int{1, 3, 2, 3, 3}, 2))
}
