package main

import "fmt"

func main() {
	var grid = [3][3]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	timer := 0
	var arr [][2]int
	var arr2 [][2]int

	if len(arr) == 0 {
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				if grid[i][j] == 1 {
					println("impossible")
					return
				}
			}
		}
		println("0")
		return
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				arr = append(arr, [2]int{i, j})
			}
		}
	}
	i := 0
	count := 0
	for true {
		if arr[i][0] != 0 && grid[arr[i][0]-1][arr[i][1]] == 1 {
			grid[arr[i][0]-1][arr[i][1]] = 2
			arr2 = append(arr2, [2]int{arr[i][0] - 1, arr[i][1]})
			count++
		}

		if arr[i][1] != 0 && grid[arr[i][0]][arr[i][1]-1] == 1 {
			grid[arr[i][0]][arr[i][1]-1] = 2
			arr2 = append(arr2, [2]int{arr[i][0], arr[i][1] - 1})
			count++
		}

		if arr[i][0] != len(grid)-1 && grid[arr[i][0]+1][arr[i][1]] == 1 {
			grid[arr[i][0]+1][arr[i][1]] = 2
			arr2 = append(arr2, [2]int{arr[i][0] + 1, arr[i][1]})
			count++
		}

		if arr[i][1] != len(grid[0])-1 && grid[arr[i][0]][arr[i][1]+1] == 1 {
			grid[arr[i][0]][arr[i][1]+1] = 2
			arr2 = append(arr2, [2]int{arr[i][0], arr[i][1] + 1})
			count++
		}
		if count == 0 {
			arr = append(arr[:i], arr[i+1:]...)
		} else {
			i++
			count = 0
		}
		if len(arr) == 0 {
			println("end: ", timer)
			break
		}
		if i == len(arr) {
			arr = append(arr, arr2...)
			arr = arr[i:]
			arr2 = arr2[:0]
			i = 0
			timer++
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				println("impossible")
				return
			}
		}
	}
	_ = timer
	fmt.Println(arr)
}
