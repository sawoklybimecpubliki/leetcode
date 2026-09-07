package main

import "fmt"

func main() {
	nums := []int{1, 2, 14, 15}
	fmt.Println(resultArray(nums))
}

func resultArray(nums []int) []int {
	l := len(nums)
	var num1, num2 []int
	num1 = append(num1, nums[0])
	num2 = append(num2, nums[1])
	for i := 2; i < l; i++ {
		if num1[len(num1)-1] > num2[len(num2)-1] {
			num1 = append(num1, nums[i])
		} else {
			num2 = append(num2, nums[i])
		}
	}
	return append(num1, num2...)
}
