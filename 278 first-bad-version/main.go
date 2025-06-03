package main

import (
	"fmt"
)

func firstBadVersion(n int) int {
	left, right := 1, n
	for left < right {
		curVersion := left + (right-left)/2
		if isBadVersion(curVersion) {
			right = curVersion
		} else {
			left = curVersion + 1
		}
	}
	return left
}

func isBadVersion(version int) bool {
	bad := 1702766719
	if version >= bad {
		return true
	}
	return false
}

func main() {
	fmt.Println(firstBadVersion(2126753390))
}
