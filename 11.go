package main

func main() {
	height := []int{8, 7, 2, 1}
	var maxHeight, square int

	for i := len(height) - 1; i > 0; i-- {

		if maxHeight < height[i] {
			maxHeight = height[i]
			for j := 0; j < i; j++ {
				if height[i] > height[j] {

					if (i-j)*height[j] > square {
						square = (i - j) * height[j]
					}

				} else {

					if (i-j)*height[i] > square {
						square = (i - j) * height[i]
					}

				}

			}
		}

	}
	println(square)
}
