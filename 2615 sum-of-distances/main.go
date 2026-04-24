package main

import "fmt"

func distance(nums []int) []int64 {
	out := make([]int64, len(nums))
	m := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = append(m[nums[i]], i)
	}
	for _, num := range m {
		l := len(num)
		if l <= 1 {
			continue
		}

		var sum int64
		for _, i := range num {
			sum += int64(i)
		}

		var prefixSum int64
		for i, idx := range num {
			idx64 := int64(idx)
			i64 := int64(i)
			m64 := int64(l)

			leftSide := i64*idx64 - prefixSum

			suffixSum := sum - prefixSum - idx64
			rightSide := suffixSum - (m64-1-i64)*idx64

			out[idx] = leftSide + rightSide

			prefixSum += idx64
		}
	}
	return out
}

func abs(num int64) int64 {
	if num > 0 {
		return num
	}
	return -num
}

func main() {
	nums := []int{1, 3, 1, 1, 2}
	fmt.Println(distance(nums))
}
