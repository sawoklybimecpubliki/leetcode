package main

import "fmt"

func main() {
	nums := []int{4, 6}
	var odd, even int
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			even++
			for j := 0; j < len(nums); j++ {
				if i != j && (nums[i]-nums[j])%2 != 0 {
					odd++
					break
				}
			}
		} else {
			for j := 0; j < len(nums); j++ {
				if i != j && (nums[i]-nums[j])%2 == 0 {
					even++
					break
				}
			}
			odd++
		}
	}
	if even == len(nums) || odd == len(nums) {
		fmt.Println(true)
		return
	}
	fmt.Println(false)
}
