package main

import "fmt"

func maxDistance(nums1 []int, nums2 []int) int {
	dis := 0
	i := 0
	j := 0
	for i < len(nums1) && j < len(nums2) {
		if i > j {
			j++
			continue
		}
		if nums1[i] <= nums2[j] {
			if dis < j-i {
				dis = j - i
			}
			j++
		} else {
			i++
		}
	}
	return dis
}

func main() {
	nums1 := []int{55, 30, 5, 4, 2}
	nums2 := []int{100, 20, 10, 10, 5}

	fmt.Println(maxDistance(nums1, nums2))
}
