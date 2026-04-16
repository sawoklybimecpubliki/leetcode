package main

import "fmt"

func solveQueries(nums []int, queries []int) []int {
	distanceLeft := make([]int, len(nums))
	distanceRight := make([]int, len(nums))
	lastSeenLeft := make(map[int]int, len(nums))
	lastSeenRight := make(map[int]int, len(nums))

	for i, j := 0, (len(nums)*2)-1; i < len(nums)*2; i, j = i+1, j-1 {
		if v, ok := lastSeenLeft[nums[i%len(nums)]]; ok {
			// If its not same element
			if v != i%len(nums) {
				distanceLeft[i%len(nums)] = i - v
				lastSeenLeft[nums[i%len(nums)]] = i
			}
		} else {
			distanceLeft[i%len(nums)] = -1
			lastSeenLeft[nums[i%len(nums)]] = i
		}
		if v, ok := lastSeenRight[nums[j%len(nums)]]; ok {
			// If its not same element
			if v != j+len(nums) {
				distanceRight[j%len(nums)] = v - j
				lastSeenRight[nums[j%len(nums)]] = j
			}
		} else {
			distanceRight[j%len(nums)] = -1
			lastSeenRight[nums[j%len(nums)]] = j
		}
	}

	res := make([]int, len(queries))
	for i, q := range queries {
		res[i] = min(distanceLeft[q], distanceRight[q])
	}
	return res
}

/*func solveQueries(nums []int, queries []int) []int {
	var out []int
	if len(nums) <= 1 {
		return []int{-1}
	}
	for _, i := range queries {
		l, r, c := i-1, i+1, 1
		for {
			if l < 0 {
				l = len(nums) - 1
			}
			if r > len(nums)-1 {
				r = 0
			}
			if nums[l] == nums[i] {
				out = append(out, c)
				break
			}
			if nums[r] == nums[i] {
				out = append(out, c)
				break
			}
			l--
			r++
			c++
			if c > len(nums)/2 {
				out = append(out, -1)
				break
			}
		}
	}
	return out
}*/

func main() {
	fmt.Println(solveQueries(
		[]int{17, 17, 20, 16, 17, 17, 6, 11, 11, 15, 17, 15, 16, 20, 20, 9, 15, 10},
		[]int{13, 15, 11, 4, 5, 3, 14, 7, 8, 1, 16, 17}))
}
