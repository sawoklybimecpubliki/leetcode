package main

import "fmt"

func findClosest(x int, y int, z int) int {
	if abs(x, z) > abs(y, z) {
		return 2
	} else if abs(x, z) < abs(y, z) {
		return 1
	}
	return 0
}

func abs(x, y int) int {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}

func main() {
	fmt.Println(findClosest(2, 5, 6))
}
