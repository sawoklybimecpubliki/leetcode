package main

func maxProfit(prices []int) int {
	profit := 0
	min := prices[0]
	for _, num := range prices {
		if num < min {
			min = num
		} else if num-min > profit {
			profit = num - min
		}
	}
	return profit
}

func main() {
	println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
