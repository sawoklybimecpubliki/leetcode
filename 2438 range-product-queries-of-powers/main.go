package main

import "fmt"

func productQueries(n int, queries [][]int) []int {
	const MOD = 1000000007
	prefix := make([]int, 0)
	prefix = append(prefix, 0)
	sum, maxexponent := 0, 0
	for i := 0; i < 32; i++ {
		if (n & (1 << i)) > 0 {
			sum += i
			prefix = append(prefix, sum)
			maxexponent = max(maxexponent, sum)
		}
	}
	cache := []int{1}
	for i := 1; i <= maxexponent; i++ {
		cache = append(cache, (cache[len(cache)-1]*2)%MOD)
	}
	ans := make([]int, 0)
	for _, query := range queries {
		left, right := query[0], query[1]
		ans = append(ans, cache[prefix[right+1]-prefix[left]])
	}
	return ans
}

func main() {
	fmt.Println(productQueries(2, [][]int{{0, 0}}))
}
