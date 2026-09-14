package main

import "fmt"

func main() {
	rec1 := []int{0, 0, 1, 1}
	rec2 := []int{1, 0, 2, 1}
	fmt.Println(isRectangleOverlap(rec1, rec2))
}

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return !(rec1[2] <= rec2[0] ||
		rec1[0] >= rec2[2] ||
		rec1[3] <= rec2[1] ||
		rec1[1] >= rec2[3])
}
