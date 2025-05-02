package main

func main() {
	grid := [][]int{{1, 3, 1, 15}, {1, 3, 3, 1}}
	sum1 := -grid[0][0]
	sum2 := -grid[1][len(grid[1])-1]
	index := 0
	for i := 0; i < len(grid[0]); i++ {
		sum1 += grid[0][i]
		sum2 += grid[1][i]
	}
	sum_max := sum2
	for i := 0; i < len(grid[0])-1; i++ {
		if sum2+grid[0][i+1]-grid[1][i] > sum2 {
			sum_max = sum2 + grid[0][i+1] - grid[1][i]
			index = i + 1
		}
		sum2 += grid[0][i+1] - grid[1][i]
	}
	for i := 0; i < index; i++ {
		grid[0][i] = 0
		grid[1][i+index] = 0
		sum1 += grid[0][i]
		sum2 += grid[1][i]
	}
	println("answer:", index, sum_max)
}
