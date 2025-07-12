package main

import (
	"container/heap"
	"fmt"
	"slices"
)

func maxEvents(events [][]int) int {
	slices.SortFunc(events, func(a, b []int) int {
		return a[0] - b[0]
	})

	// The end times of ongoing but not-yet-attended events
	available := &intHeap{}
	day := 0
	result := 0

	for 0 < len(events) || 0 < available.Len() {
		for 0 < available.Len() && (*available)[0] < day {
			heap.Pop(available)
		}
		if available.Len() == 0 && 0 < len(events) {
			day = events[0][0]
		}

		for 0 < len(events) && events[0][0] <= day {
			heap.Push(available, events[0][1])
			events = events[1:]
		}

		if available.Len() == 0 {
			break
		}
		heap.Pop(available)
		day++
		result++
	}
	return result
}

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	fmt.Println(maxEvents([][]int{{1, 2}, {2, 2}, {3, 3}, {3, 4}, {3, 4}}))
	//{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}, {1, 7}
}
