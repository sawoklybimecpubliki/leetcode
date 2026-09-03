package main

import "fmt"

func main() {
	nums := []int{4, 6}
	fmt.Println(uniformArray(nums))
}

const mx = 1000000001

func uniformArray(nums1 []int) bool {
	minOdd, minEven := mx, mx
	for _, n := range nums1 {
		if n%2 == 0 {
			minEven = min(minEven, n)
		} else {
			minOdd = min(minOdd, n)
		}
	}
	return minEven == mx || minOdd == mx || minEven > minOdd
}
