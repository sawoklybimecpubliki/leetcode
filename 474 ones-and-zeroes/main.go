package main

import "fmt"

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, s := range strs {
		zeros, ones := countZeroOne(s)
		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				if dp[i-zeros][j-ones]+1 > dp[i][j] {
					dp[i][j] = dp[i-zeros][j-ones] + 1
				}
			}
		}
	}
	return dp[m][n]
}

func countZeroOne(s string) (int, int) {
	zeros, ones := 0, 0
	for _, ch := range s {
		if ch == '0' {
			zeros++
		} else {
			ones++
		}
	}
	return zeros, ones
}

func main() {
	fmt.Println(findMaxForm([]string{"10", "0", "1"}, 1, 1))
}
