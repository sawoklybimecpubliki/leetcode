package main

import (
	"fmt"
)

func compareVersion(version1 string, version2 string) int {
	index1, index2 := 0, 0
	length1, length2 := len(version1), len(version2)

	for index1 < length1 || index2 < length2 {
		number1 := 0
		for index1 < length1 && version1[index1] != '.' {
			number1 = number1*10 + int(version1[index1]-'0')
			index1++
		}
		index1++

		number2 := 0
		for index2 < length2 && version2[index2] != '.' {
			number2 = number2*10 + int(version2[index2]-'0')
			index2++
		}
		index2++

		if number1 > number2 {
			return 1
		} else if number1 < number2 {
			return -1
		}
	}
	return 0
}

func main() {
	fmt.Println(compareVersion("1.2", "1.10"))
}
