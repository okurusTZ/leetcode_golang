package main

import "fmt"

// 1 1 1
// 1 2 0
// 1 0 1

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	color := image[sr][sc]
	if newColor == color {
		return image
	}

	dfs(image, sr, sc, newColor, color)

	return image
}

func dfs(image [][]int, sr, sc, newColor, oldCocolor int) {
	if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) || image[sr][sc] != oldCocolor {
		return
	}
	image[sr][sc] = newColor
	dfs(image, sr-1, sc, newColor, oldCocolor)
	dfs(image, sr+1, sc, newColor, oldCocolor)
	dfs(image, sr, sc+1, newColor, oldCocolor)
	dfs(image, sr, sc-1, newColor, oldCocolor)
}

func main() {
	fmt.Println(floodFill([][]int{
		{1, 1, 0},
		{1, 1, 0},
		{1, 0, 1},
	}, 1, 1, 2))
}
