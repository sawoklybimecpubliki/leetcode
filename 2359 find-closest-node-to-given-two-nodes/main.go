package main

import "fmt"

func getWeight(edges []int, start int) []int {

	n := len(edges)
	weight := make([]int, n)
	for i := range weight {
		weight[i] = -1
	}
	d := 0
	for start != -1 && weight[start] == -1 {
		weight[start] = d
		d++
		start = edges[start]
	}
	return weight
}

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	weight1 := getWeight(edges, node1)
	weight2 := getWeight(edges, node2)
	result := -1
	minweight := len(edges)
	for i := 0; i < len(weight1); i++ {
		if weight1[i] != -1 && weight2[i] != -1 {
			maxweight := weight1[i]
			if weight2[i] > maxweight {
				maxweight = weight2[i]
			}
			if maxweight < minweight {
				minweight = maxweight
				result = i
			}
		}
	}

	return result
}

func main() {
	fmt.Println(closestMeetingNode([]int{2, 2, 3, -1}, 0, 1))
}
