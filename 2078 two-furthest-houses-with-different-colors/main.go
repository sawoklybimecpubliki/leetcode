package main

import "fmt"

func maxDistance(colors []int) int {
	var distance int
	n := len(colors)
	for i := n - 1; i >= 0; i-- {
		if colors[0] != colors[i] && distance < i {
			distance = i
		}
	}
	for i := 0; i < n; i++ {
		if colors[n-1] != colors[i] && distance < n-1-i {
			distance = n - 1 - i
		}
	}
	return distance
}

/*
	for i := 0; i < len(colors)-1; i++ {
		for j := i + 1; j < len(colors); j++ {
			if colors[i] != colors[j] {
				if distance < j-i {
					distance = j - i
				}
			}
		}
	}
*/
func main() {
	colors := []int{4, 4, 4, 11, 4, 4, 11, 4, 4, 4, 4, 4}
	fmt.Println(maxDistance(colors))
}
