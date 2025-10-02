package main

import "fmt"

func maxBottlesDrunk(numBottles int, numExchange int) int {
	result := numBottles
	emptyBottles := numBottles
	for numBottles > 0 {
		numBottles = 0
		for emptyBottles >= numExchange {
			numBottles++
			emptyBottles -= numExchange
			numExchange++
		}
		emptyBottles += numBottles
		result += numBottles
	}

	return result
}

func main() {
	fmt.Println(maxBottlesDrunk(10, 3))
}
