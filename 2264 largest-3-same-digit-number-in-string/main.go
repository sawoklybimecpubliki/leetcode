package main

import "fmt"

func largestGoodInteger(num string) string {
	best := ""
	for i := 0; i+2 < len(num); i++ {
		if num[i] == num[i+1] && num[i] == num[i+2] {
			cur := num[i : i+3]
			if cur > best {
				best = cur
			}
		}
	}
	return best
}

func main() {
	fmt.Println(largestGoodInteger("6777133339"))
}
