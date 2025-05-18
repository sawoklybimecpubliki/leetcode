package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	var dist = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var stack [][]int

	curColor := image[sr][sc]
	image[sr][sc] = color

	for i := 0; i < 4; i++ {

		if sr+dist[i][0] < 0 || sr+dist[i][0] > len(image)-1 ||
			sc+dist[i][1] < 0 || sc+dist[i][1] > len(image[0])-1 {
			continue
		}
		if image[sr+dist[i][0]][sc+dist[i][1]] == curColor {
			stack = append(stack, []int{sr + dist[i][0], sc + dist[i][1]})
		}
	}
	j := 0
	for {
		if len(stack) == 0 {
			break
		}
		x := stack[j][0]
		y := stack[j][1]
		if image[x][y] == curColor {
			image[x][y] = color
		}
		for i := 0; i < 4; i++ {
			if x+dist[i][0] < 0 || x+dist[i][0] > len(image)-1 ||
				y+dist[i][1] < 0 || y+dist[i][1] > len(image[0])-1 {
				continue
			}
			if image[x+dist[i][0]][y+dist[i][1]] == curColor {
				stack = append(stack, []int{x + dist[i][0], y + dist[i][1]})
			}
		}
		j++
		if j >= len(stack) || j > len(image)+len(image[0]) {
			break
		}
	}

	return image

}

func main() {
	fmt.Println(floodFill(
		[][]int{{0, 0, 0}, {0, 0, 0}},
		0, 0, 0))
}
