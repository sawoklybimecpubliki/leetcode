package main

import "fmt"

func removeElement(nums []int, val int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			result++
		}
	}
	/*
		   r := 0
			for i := 0; i < len(nums); i++ {
				if nums[i] != val {
					nums[r] = nums[i]
		            r++
				}
			}
			return r
	*/
	return result
}

func main() {
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
