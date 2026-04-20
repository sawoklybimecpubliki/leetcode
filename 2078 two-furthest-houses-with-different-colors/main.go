package main

import "fmt"

func maxDistance(colors []int) int {
	var distance int
	for i := 0; i < len(colors)-1; i++ {
		for j := i + 1; j < len(colors); j++ {
			if colors[i] != colors[j] {
				if distance < j-i {
					distance = j - i
				}
			}
		}
	}
	return distance
}

func main() {
	colors := []int{1, 8, 3, 8, 3}
	fmt.Println(maxDistance(colors))
}
