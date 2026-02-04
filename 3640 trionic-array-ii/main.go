package main

import "fmt"

func maxSumTrionic(nums []int) int64 {

	n := len(nums)
	var res int64 = -10000000000000000

	for i := 1; i < n-2; i++ {

		a, b := i, i
		net := int64(nums[a])

		for b+1 < n && nums[b+1] < nums[b] {
			net += int64(nums[b+1])
			b++
		}

		if a == b {
			continue
		}

		c := b

		var left, right int64 = 0, 0
		var lx, rx int64 = -1 << 63, -1 << 63

		for a-1 >= 0 && nums[a-1] < nums[a] {
			left += int64(nums[a-1])
			if left > lx {
				lx = left
			}
			a--
		}

		if a == i {
			continue
		}

		for b+1 < n && nums[b+1] > nums[b] {
			right += int64(nums[b+1])
			if right > rx {
				rx = right
			}
			b++
		}

		if b == c {
			continue
		}

		if lx+net+rx > res {
			res = lx + net + rx
		}

		i = b - 1
	}

	return res
}

func main() {
	fmt.Println(maxSumTrionic([]int{0, -2, -1, -3, 0, 2, -1}))
}
