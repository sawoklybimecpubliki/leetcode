package main

import "fmt"

func insert(intervals [][]int, newIntervals []int) [][]int {

	minMaxInterval := []int{-1, -1}
	var outInterval [][]int
	i := 0
	if len(intervals) == 0 {
		return append(intervals, newIntervals)
	}
	for i < len(intervals) && newIntervals[0] > intervals[i][0] {
		if newIntervals[0] > intervals[i][1] {
			outInterval = append(outInterval, intervals[i])
			i++
		} else {
			minMaxInterval[0] = intervals[i][0]
			break
		}
	}
	if minMaxInterval[0] == -1 {
		minMaxInterval[0] = newIntervals[0]
	}
	for i < len(intervals) && newIntervals[1] >= intervals[i][0] {
		if newIntervals[1] > intervals[i][1] {
			i++
		} else if newIntervals[1] <= intervals[i][1] {
			minMaxInterval[1] = intervals[i][1]
			i++
			break
		}
	}
	if minMaxInterval[1] == -1 {
		minMaxInterval[1] = newIntervals[1]
	}
	outInterval = append(outInterval, minMaxInterval[:])
	outInterval = append(outInterval, intervals[i:]...)
	return outInterval
}

func main() {
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
}
