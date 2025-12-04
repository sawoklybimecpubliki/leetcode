package main

import "fmt"

func countCollisions(directions string) int {
	i, j := 0, len(directions)-1
	for i < len(directions) && directions[i] == 'L' {
		i++
	}
	for j >= 0 && directions[j] == 'R' {
		j--
	}
	count := 0
	for k := i; k <= j; k++ {
		if directions[k] != 'S' {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countCollisions("RLRSLL"))
}
