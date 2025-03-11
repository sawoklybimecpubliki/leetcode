package main

import "fmt"

func recoursive(nums []int, out []int, result *[][]int) {
	for i, s := range nums {
		out = append(out, s)
		tmp := make([]int, len(out))
		copy(tmp, out)
		*result = append(*result, tmp)
		recoursive(nums[i+1:], out, result)
		out = out[:len(out)-1]
	}
}

func subsets(nums []int) [][]int {
	var result = [][]int{{}}
	var out []int
	recoursive(nums[:], out, &result)
	return result
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
