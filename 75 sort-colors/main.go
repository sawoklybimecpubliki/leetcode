package main

import "fmt"

func sortColors(nums []int) {
	fmt.Println(nums)
	for c := 0; c < len(nums)/2+1; c++ {
		j := -1
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				if j != -1 {
					nums[j], nums[i+1] = nums[i+1], nums[j]
					j = i
				} else {
					nums[i], nums[i+1] = nums[i+1], nums[i]
				}
			} else if nums[i] == nums[i+1] {
				j = i
			} else {
				j = -1
			}
			fmt.Println(nums)
		}
	}

}

func main() {
	sortColors([]int{1, 2, 2, 2, 2, 0, 0, 0, 1, 1})
}
