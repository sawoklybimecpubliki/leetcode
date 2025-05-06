package main

import "fmt"

func numTilings(n int) int {
	const MOD = 1000000007
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	a, b, c := 1, 1, 2
	for i := 3; i <= n; i++ {
		d := (2*c + a) % MOD
		a, b, c = b, c, d
	}
	return c
}

func main() {
	fmt.Println(numTilings(3))
}
