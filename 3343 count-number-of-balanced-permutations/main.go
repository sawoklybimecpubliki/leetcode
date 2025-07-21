package main

import (
	"fmt"
	"strconv"
)

func countBalancedPermutations(num string) int {
	const mod = 1000000007
	var result int
	var arr []int
	var cache = make(map[string][]int)
	for _, i := range num {
		n, _ := strconv.Atoi(string(i))
		arr = append(arr, n)
	}
	permute(arr, 0, cache)
	for _, i := range cache {
		sum_odd, sum_even := sum(i)
		if sum_odd == sum_even {
			result++
		}
	}

	fmt.Println(cache)
	return result
}

func permute(nums []int, index int, c map[string][]int) {
	if index >= len(nums)-1 { // Базовый случай: последняя позиция
		var str string
		var arr []int
		arr = append(arr, nums...)
		for _, i := range arr {
			str += strconv.Itoa(i)
		}
		if _, ok := c[str]; !ok {
			c[str] = arr
		}
		return
	}

	for i := index; i < len(nums); i++ {
		nums[index], nums[i] = nums[i], nums[index] // Меняем местами элементы
		permute(nums, index+1, c)                   // Рекурсия для следующего индекса
		nums[index], nums[i] = nums[i], nums[index] // Возвращаем обратно (backtracking)
	}
}

func sum(arr []int) (int, int) {
	sum_odd, sum_even := 0, 0
	l := 0
	if len(arr)%2 == 0 {
		l = len(arr) / 2
	} else {
		l = (len(arr) - 1) / 2
	}
	for i, j, c := 1, 0, 0; ; c++ {
		if c < l {
			sum_odd += arr[i]
		}

		if j < len(arr) {
			sum_even += arr[j]
		} else {
			break
		}
		i += 2
		j += 2
	}
	return sum_odd, sum_even
}

func main() {
	fmt.Println("answer: ", countBalancedPermutations("875654570"))
}
