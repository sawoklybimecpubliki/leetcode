package main

import "fmt"

func dfs(grid [][]int, point []int) {
	var dir = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m := len(grid)
	n := len(grid[0])

	if point[0] < 0 || point[0] >= m || point[1] < 0 || point[1] >= n || grid[point[0]][point[1]] != 1 {
		return
	}

	grid[point[0]][point[1]] = 0
	for i := 0; i < 4; i++ {
		dfs(grid, []int{point[0] + dir[i][0], point[1] + dir[i][1]})
	}

	return
}

func numEnclaves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for _, i := range []int{0, m - 1} {
		for j := 0; j < n; j++ {

			if grid[i][j] == 1 {
				dfs(grid, []int{i, j})
			}

		}
	}
	for _, i := range []int{0, n - 1} {
		for j := 0; j < m; j++ {

			if grid[j][i] == 1 {
				dfs(grid, []int{j, i})
			}

		}
	}
	count := 0
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 1 {
				count++
			}
		}
	}
	return count
}

func main() {
	fmt.Println(numEnclaves([][]int{
		{0, 0, 0, 1, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 0, 0, 1, 0, 1, 1, 1},
		{0, 0, 0, 1, 1, 1, 0, 1, 0, 0},
		{0, 1, 1, 0, 0, 0, 1, 0, 1, 0},
		{0, 1, 1, 1, 1, 1, 0, 0, 1, 0},
		{0, 0, 1, 0, 1, 1, 1, 1, 0, 1},
		{0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 1, 0, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0, 0, 1}}))
}

/*
package main

import "fmt"

func dfs(grid [][]int, x, y int, seenBorder bool) (int, bool) {
  var dir = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
  m := len(grid)
  n := len(grid[0])

  if x < 0 || y < 0 || x >= m || y >= n || grid[x][y] != 1 {
    return 0, false
  }

  grid[x][y] = 2
  val := 1
  isBorder := x == 0 || y == 0 || x == m-1 || y == n-1 || seenBorder

  for i := 0; i < 4; i++ {
    subval, subIsBorder := dfs(grid, x+dir[i][0], y+dir[i][1], isBorder)
    isBorder = isBorder || subIsBorder
    val += subval
  }

  return val, isBorder
}

func numEnclaves(grid [][]int) int {
  m := len(grid)
  n := len(grid[0])
  ans := 0
  for i := 0; i < m; i++ {
    for j := 0; j < n; j++ {
      if grid[i][j] == 1 {
        if t, v := dfs(grid, i, j, false); !v {
          ans += t
        }
      }
    }
  }

  return ans
}

*/
