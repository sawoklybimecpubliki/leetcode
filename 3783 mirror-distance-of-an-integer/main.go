package main

import "fmt"

func mirrorDistance(n int) int {
	r := reverse(n)
	if n > r {
		return n - r
	}
	return r - n
}

func reverse(n int) int {
	var rev int
	for n%10 != n {
		rev = rev*10 + n%10
		n = n / 10
	}
	rev = rev*10 + n%10
	return rev
}

func main() {
	fmt.Println(mirrorDistance(25))
}
