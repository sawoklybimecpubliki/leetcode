package main

import "fmt"

func minSum(nums1 []int, nums2 []int) int64 {
	sum1, sum2 := 0, 0
	count1, count2 := 0, 0
	l1, l2 := len(nums1), len(nums2)
	for i := 0; ; {
		if i >= l1 && i >= l2 {
			break
		}

		if i < l1 {
			sum1 += nums1[i]
			if nums1[i] == 0 {
				count1++
			}
		}

		if i < l2 {
			sum2 += nums2[i]
			if nums2[i] == 0 {
				count2++
			}
		}
		i++
	}

	var result int
	var add [2]int

	for add[0], add[1] = count1, count2; ; {
		if sum1+add[0] == sum2+add[1] {
			result = sum1 + add[0]
			break
		} else if sum1+add[0] < sum2+add[1] {
			if count1 == 0 {
				return -1
			}
			add[0] = sum2 + add[1] - sum1

		} else if count2 != 0 {
			add[1] = sum1 + add[0] - sum2
		} else {
			return -1
		}

	}
	return int64(result)
}

func main() {
	fmt.Println(minSum(
		[]int{2, 0, 2, 0},
		[]int{1, 4}))
}
