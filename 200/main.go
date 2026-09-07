package main

import "fmt"

var (
	nums = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	res  = 200
)

func main() {
	dfs(1, 0, nums[0], fmt.Sprint(nums[0]))
}

func dfs(i, sum, curr int, strOut string) {
	if i == len(nums) {
		if sum+curr == res {
			fmt.Printf("%s = 200\n", strOut)
		}
		return
	}
	num := nums[i]
	newCurr := curr*10 + num
	dfs(i+1, sum, newCurr, strOut+fmt.Sprint(num))
	dfs(i+1, sum+curr, num, strOut+"+"+fmt.Sprint(num))
	dfs(i+1, sum+curr, -num, strOut+"-"+fmt.Sprint(num))
}
