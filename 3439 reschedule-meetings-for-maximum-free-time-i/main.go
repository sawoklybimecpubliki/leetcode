package main

import "fmt"

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	count := len(startTime)
	maxFree := 0
	prefixSum := make([]int, count+1)

	for i := 0; i < count; i++ {
		prefixSum[i+1] = prefixSum[i] + (endTime[i] - startTime[i])
	}

	for i := k - 1; i < count; i++ {
		occupied := prefixSum[i+1] - prefixSum[i+1-k]
		windowEnd := eventTime
		if i != count-1 {
			windowEnd = startTime[i+1]
		}
		windowStart := 0
		if i != k-1 {
			windowStart = endTime[i-k]
		}
		freeTime := windowEnd - windowStart - occupied
		if freeTime > maxFree {
			maxFree = freeTime
		}
	}

	return maxFree
}

func main() {
	fmt.Println(maxFreeTime(5, 2, []int{0, 1, 2, 3, 4}, []int{1, 2, 3, 4, 5}))
}
