package main

import "fmt"

func minSumOfLengths(arr []int, target int) int {
	n := len(arr)
	best, answer := n+1, n+1
	start, sum := 0, 0
	prefixStart, prefixEnd, prefixSum := 0, 0, 0

	for end, value := range arr {
		sum += value
		for sum > target {
			sum -= arr[start]
			start++
		}
		if sum != target {
			continue
		}

		// Only process earlier windows ending strictly before start.
		for prefixEnd < start {
			prefixSum += arr[prefixEnd]
			for prefixSum > target {
				prefixSum -= arr[prefixStart]
				prefixStart++
			}
			if prefixSum == target {
				length := prefixEnd - prefixStart + 1
				if length < best {
					best = length
				}
			}
			prefixEnd++
		}

		if best <= n {
			total := best + end - start + 1
			if total < answer {
				answer = total
			}
		}
	}

	if answer == n+1 {
		return -1
	}
	return answer
}

func main() {
	fmt.Println(minSumOfLengths([]int{3, 2, 2, 4, 3}, 3))
}
