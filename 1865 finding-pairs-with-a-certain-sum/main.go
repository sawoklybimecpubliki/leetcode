package main

import "fmt"

type FindSumPairs struct {
	nums1 []int
	nums2 []int
	freq  map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	freq := make(map[int]int)
	for _, num := range nums2 {
		freq[num]++
	}
	return FindSumPairs{nums1: nums1, nums2: nums2, freq: freq}
}

func (this *FindSumPairs) Add(index int, val int) {
	/*	tmp := make([]int, index, cap(this.nums2))
		copy(tmp, this.nums2[:index])
		tmp = append(tmp, val)
		tmp = append(tmp, this.nums2[index:]...)
		this.nums2 = tmp*/
	this.freq[this.nums2[index]]--
	this.nums2[index] += val
	this.freq[this.nums2[index]]++
}

func (this *FindSumPairs) Count(tot int) int {
	count := 0
	for _, num := range this.nums1 {
		count += this.freq[tot-num]
	}
	return count
}

func main() {
	obj := Constructor([]int{1, 1, 2, 2, 2, 3}, []int{1, 4, 5, 2, 5, 4})
	fmt.Println(obj.Count(7))
	obj.Add(3, 2)
	fmt.Println(obj.Count(8))
	fmt.Println(obj.Count(4))
	obj.Add(0, 1)
	obj.Add(1, 1)
	fmt.Println(obj.Count(7))
}

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */
