package main

import "fmt"

func lexicalOrder(n int) []int {
	var out []int
	cur := 1
	for i := 0; i < n; i++ {
		out = append(out, cur)
		if cur*10 <= n {
			cur *= 10
		} else {
			for cur%10 == 9 || cur >= n {
				cur /= 10
			}
			cur++
		}
	}
	return out
}

func main() {
	fmt.Println(lexicalOrder(1012))
}
