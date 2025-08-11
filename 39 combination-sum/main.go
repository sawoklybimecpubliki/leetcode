package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var out []int
	var result [][]int
	sum(candidates, target, out, &result)
	fmt.Println(result)
	return [][]int{}
}

func sum(nums []int, target int, out []int, result *[][]int) {
	for i, num := range nums {
		if target-num > 0 {
			out = append(out, num)
			sum(nums[i:], target-num, out, result)
		} else if target-num < 0 {
			sum(nums[i+1:], target, out, result)
			out = out[:len(out)]
			return
		} else {
			out = append(out, num)
			tmp := make([]int, len(out))
			copy(tmp, out)
			*result = append(*result, tmp)
			fmt.Println("OUT: ", result)

		}
		out = out[:len(out)-1]
	}
}

func main() {
	combinationSum([]int{4, 2, 8}, 8)
}
