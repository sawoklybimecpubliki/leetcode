package main

import "fmt"

func countSubarrays(nums []int, k int64) int64 {
	var sum, result int
	j := 0
	leng := len(nums)
	for i := 0; i < leng; i++ {
		sum += nums[i]
		for sum*(i-j+1) >= int(k) {
			sum -= nums[j]
			j++
		}
		result += (i - j + 1)
	}

	return int64(result)
}

func main() {
	fmt.Println(countSubarrays([]int{9, 5, 3, 8, 4, 7, 2, 7, 4, 5, 4, 9, 1, 4, 8, 10, 8, 10, 4, 7}, 4))
}
