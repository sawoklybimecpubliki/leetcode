package main

import "fmt"

func minMirrorPairDistance(nums []int) int {
	distance := len(nums)
	m := make(map[int]int)
	for i, num := range nums {
		if prevIndex, ok := m[num]; ok {
			if i-prevIndex < distance {
				distance = i - prevIndex
			}
		}
		rev := reverse(num)
		m[rev] = i
	}
	if distance == len(nums) {
		return -1
	}
	return distance
}

func main() {
	nums := []int{120, 21}
	fmt.Println(minMirrorPairDistance(nums))
}

func reverse(num int) int {
	rev := 0
	for {
		rev *= 10
		if num%10 == num {
			rev += num
			break
		}
		rev += num % 10
		num /= 10
	}
	return rev
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
