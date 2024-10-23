package game

import (
	"math/rand"
)

func RandomMap(height int, width int) [][]int {
	grid := make([][]int, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]int, width)
	}
	percent := rand.Intn(70) + 10
	aliveCells := (height * width) * percent / 100
	if aliveCells < 3 {
		aliveCells = 3
	}

	for aliveCells > 0 {
		row := rand.Intn(height)
		col := rand.Intn(width)
		if grid[row][col] != 1 {
			grid[row][col] = 1
			aliveCells--
		}
	}
	return grid
}
