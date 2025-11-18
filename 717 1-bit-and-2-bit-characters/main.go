package main

import "fmt"

func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	i := 0
	for i = 0; i < n-1; i++ {
		if bits[i] == 1 {
			i++
		}
	}
	return i == n-1
}

func main() {
	fmt.Println(isOneBitCharacter([]int{0, 1, 1, 0}))
}
