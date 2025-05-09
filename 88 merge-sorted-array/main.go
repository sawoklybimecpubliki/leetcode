package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	/*result := make([]int, m+n)
	for i, j := 0, 0; i < m || j < n; {
		if i < m && nums1[i] <= nums2[j] {
			result[i+j] = nums1[i]
			i++
		} else {
			result[i+j] = nums2[j]
			j++
		}
	}
	nums1 = result*/
	for i, j := 0, 0; ; {
		if j >= n {
			break
		}
		if i == m && j != n {
			nums1 = append(nums1[:i], nums2[j:]...)
			break
		}
		if nums1[i] > nums2[j] {
			temp := append([]int{}, nums1[i:len(nums1)-1]...)
			nums1 = append(nums1[:i], nums2[j])
			nums1 = append(nums1, temp...)
			j++
			m++
		} else {
			i++
		}
	}
	return nums1
}

func main() {
	fmt.Println(merge([]int{4, 0, 0, 0, 0, 0}, 1, []int{1, 2, 3, 5, 6}, 5))
}
