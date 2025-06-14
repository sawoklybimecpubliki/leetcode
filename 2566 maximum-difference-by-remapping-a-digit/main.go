package main

import "fmt"

func minMaxDifference(num int) int {
	var nums, out_max, out_min []int
	for i := 1; i <= num; i *= 10 {
		nums = append(nums, (num%(i*10))/i)
	}
	for i, n, m := len(nums)-1, -1, nums[len(nums)-1]; i >= 0; i-- {
		if nums[i] != 9 && n == -1 {
			n = nums[i]
		}
		if nums[i] == m {
			out_min = append(out_min, 0)
		} else {
			out_min = append(out_min, nums[i])
		}
		if nums[i] == n {
			out_max = append(out_max, 9)
		} else {
			out_max = append(out_max, nums[i])
		}
	}
	max_val, min_val := 0, 0
	pow := 1
	for i := len(nums) - 1; i >= 0; i-- {
		max_val += out_max[i] * (pow)
		min_val += out_min[i] * (pow)
		pow *= 10
	}
	return max_val - min_val
}

func main() {
	fmt.Println(minMaxDifference(90693669))
}
