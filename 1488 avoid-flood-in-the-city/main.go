package main

import "fmt"

func avoidFlood(rains []int) []int {
	isFull := make(map[int]int)
	zeroIndex := make([]int, 0)
	ans := make([]int, len(rains))

	// Initialize answer
	for i := range ans {
		if rains[i] == 0 {
			ans[i] = 1
		} else {
			ans[i] = -1
		}
	}

	for i := 0; i < len(rains); i++ {
		if rains[i] == 0 {
			zeroIndex = append(zeroIndex, i)
		} else if _, ok := isFull[rains[i]]; ok {
			// Need a dry day after the last rain
			if len(zeroIndex) == 0 {
				return []int{}
			}
			zeroIdx := getMinIndex(zeroIndex, isFull[rains[i]])
			if zeroIdx == -1 {
				return []int{}
			}
			ans[zeroIndex[zeroIdx]] = rains[i]
			zeroIndex = append(zeroIndex[:zeroIdx], zeroIndex[zeroIdx+1:]...)
		}
		isFull[rains[i]] = i
	}
	return ans
}

func getMinIndex(arr []int, v int) int {
	left, right := 0, len(arr)-1
	for left < right {
		m := left + (right-left)/2
		if arr[m] >= v {
			right = m
		} else {
			left = m + 1
		}
	}
	if arr[right] < v {
		return -1
	}
	return right
}

func main() {
	fmt.Println(avoidFlood([]int{0, 1, 1}))
}
