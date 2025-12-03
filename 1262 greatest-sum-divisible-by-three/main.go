package main

import "fmt"

func maxSumDivThree(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	sumRem := sum % 3
	if sumRem == 0 {
		return sum
	}

	rem1 := sumRem
	minRem1Num := sum
	rem2 := 3 - sumRem
	minRem2Nums := []int{sum, sum}

	for _, num := range nums {
		switch num % 3 {
		case rem1:
			minRem1Num = min(minRem1Num, num)
		case rem2:
			if num < minRem2Nums[0] {
				minRem2Nums[1] = minRem2Nums[0]
				minRem2Nums[0] = num
			} else if num < minRem2Nums[1] {
				minRem2Nums[1] = num
			}
		}
	}
	return sum - min(minRem1Num, minRem2Nums[0]+minRem2Nums[1])
}

func main() {
	fmt.Println(maxSumDivThree([]int{3, 6, 5, 1, 8}))
}
