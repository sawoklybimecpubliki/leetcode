package main

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+i, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func permutations(permut []int, nums []int, result *[][]int) {
	for i := 0; i < len(nums); i++ {
		permut = append(permut, nums[0])
		nums[0], nums[i] = nums[i], nums[0]
		permutations(permut[:], nums[1:], result)
		nums[0], nums[i] = nums[i], nums[0]
		permut = permut[:len(permut)-1]
	}
	if len(nums) == 0 {
		arr := make([]int, len(permut))
		copy(arr, permut)
		*result = append(*result, arr)
	}
	return
}

func main() {
	// с помощью рекурсии переходя по массиву in range
	nums := []int{1, 2, 3, 4}
	var out [][]int
	for i := 0; i < len(nums); i++ {
		nums[0], nums[i] = nums[i], nums[0]
		permutations(nums[0:1], nums[1:], &out)
		nums[0], nums[i] = nums[i], nums[0]
	}
	fmt.Println("len = ", len(out))
	fmt.Println("n = ", factorial(len(nums)))
	fmt.Println(out)
	/*
		for i := 0; i < len(nums); i++ {
			tmpSlice := make([]int, len(nums), cap(nums))
			copy(tmpSlice, nums)
			out = append(out, tmpSlice)

			for numIndex := len(nums) - 2; numIndex > 0; numIndex-- {
				tmpSlice := make([]int, len(nums), cap(nums))
				copy(tmpSlice, nums)
				reverse(nums[numIndex:])
				fmt.Println("first: ", nums)
				//fmt.Printf("address of slice nums %p \n", &nums)
				//fmt.Printf("address of slice tmp %p \n", &tmpSlice)
				copy(tmpSlice, nums)
				out = append(out, tmpSlice)
				reverse(nums[numIndex:])
				fmt.Println("second: ", nums)
			}

			reverse(nums[:])
		}*/
}
