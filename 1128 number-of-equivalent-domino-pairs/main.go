package main

import "fmt"

func numEquivDominoPairs(dominoes [][]int) int {
	var result int
	count := make(map[[2]int]int)
	for i := 0; i < len(dominoes); i++ {
		if dominoes[i][0] < dominoes[i][1] {
			count[[2]int{dominoes[i][0], dominoes[i][1]}]++
		} else {
			count[[2]int{dominoes[i][1], dominoes[i][0]}]++
		}

	}
	for _, i := range count {
		fmt.Println(i)
		if i > 1 {
			for j := range i {
				result += j
			}
		}
	}
	return result
}

func main() {
	fmt.Println(numEquivDominoPairs([][]int{
		{1, 2}, {2, 1},
		{1, 2}, {1, 1}, {1, 2}}))
}
