package main

import "fmt"

func numWaterBottles(numBottles int, numExchange int) int {
	result := numBottles
	emptyBottles := numBottles
	for numBottles > 0 {
		numBottles = emptyBottles / numExchange
		emptyBottles = emptyBottles + numBottles - (numBottles * numExchange)
		result += numBottles
	}
	return result
}

func main() {
	fmt.Println(numWaterBottles(9, 3))
}
