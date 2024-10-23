package game

import (
	. "crunch03/messages"
	"fmt"
	//"strconv"
	. "crunch03/colors"
	"time"
)

func Run(height int, width int, grid [][]int, options map[string]string) {
	modes := TurnModes(options)
	delay := SetDelay(options)

	tick := 1
	aliveCells := StartNumberOfAliveCells(height, width, grid)
	for {
		if aliveCells == 0 {
			StopMessage()
			return
		}
		ClearScreen()
		VerboseMode(modes["--verbose"], tick, height, width, aliveCells, delay)
		PrintGrid(height, width, grid, modes["--colored"])
		grid, aliveCells = NextGeneration(modes["--footprints"], modes["--edges-portal"], height, width, grid)
		tick++
		Delay(time.Duration(delay))
	}
}

func StartNumberOfAliveCells(height int, width int, grid [][]int) int {
	cnt := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == 1 {
				cnt++
			}
		}
	}
	return cnt
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func Delay(period time.Duration) {
	time.Sleep(period * time.Millisecond)
}

func PrintGrid(height int, width int, grid [][]int, color bool) {
	dead := ". "
	live := "× "
	past := "∘ "
	if color {
		dead = Color(". ", "blue", "")
		live = Color("× ", "green", "")
		past = Color("∘ ", "ocean", "")
	}
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if grid[row][col] == 0 {
				fmt.Print(dead)
			} else if grid[row][col] == 1 {
				fmt.Print(live)
			} else {
				fmt.Printf(past)
			}
		}
		fmt.Println()
	}
}

func NextGeneration(footprints bool, portal bool, height int, width int, grid [][]int) ([][]int, int) {
	newGrid := make([][]int, height)
	for row := 0; row < height; row++ {
		newGrid[row] = make([]int, width)
	}
	numberAliveCells := 0
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			aliveCells := CountAliveCells(portal, row, col, height, width, grid)
			if footprints {
				if grid[row][col] == 0 || grid[row][col] == 2 {
					if aliveCells == 3 {
						newGrid[row][col] = 1
						numberAliveCells++
					} else {
						newGrid[row][col] = grid[row][col]
					}
				}
				if grid[row][col] == 1 {
					if aliveCells < 2 {
						newGrid[row][col] = 2
					} else if aliveCells == 2 || aliveCells == 3 {
						newGrid[row][col] = 1
						numberAliveCells++
					} else if aliveCells > 3 {
						newGrid[row][col] = 2
					}
				}
			} else {
				if grid[row][col] == 0 {
					if aliveCells == 3 {
						newGrid[row][col] = 1
						numberAliveCells++
					} else {
						newGrid[row][col] = 0
					}
				}
				if grid[row][col] == 1 {
					if aliveCells < 2 {
						newGrid[row][col] = 0
					} else if aliveCells == 2 || aliveCells == 3 {
						newGrid[row][col] = 1
						numberAliveCells++
					} else if aliveCells > 3 {
						newGrid[row][col] = 0
					}
				}
			}
		}
	}
	return newGrid, numberAliveCells
}

func CountAliveCells(portal bool, row int, col int, height int, width int, grid [][]int) int {
	count := 0
	count += CheckCellRange(portal, row-1, col, height, width, grid)   // Up
	count += CheckCellRange(portal, row+1, col, height, width, grid)   // Down
	count += CheckCellRange(portal, row, col-1, height, width, grid)   // Left
	count += CheckCellRange(portal, row, col+1, height, width, grid)   // Right
	count += CheckCellRange(portal, row-1, col-1, height, width, grid) // Up-left
	count += CheckCellRange(portal, row-1, col+1, height, width, grid) // Up-right
	count += CheckCellRange(portal, row+1, col-1, height, width, grid) // Down-left
	count += CheckCellRange(portal, row+1, col+1, height, width, grid) // Down-right
	return count
}

func CheckCellRange(portal bool, row int, col int, height int, width int, grid [][]int) int {
	if portal {
		if row < 0 {
			row = height - 1
		}
		if row >= height {
			row = 0
		}
		if col < 0 {
			col = width
		}
		if col >= width {
			col = 0
		}
		if grid[row][col] == 1 {
			return 1
		} else {
			return 0
		}
	} else {
		if row < 0 || col < 0 || row >= height || col >= width {
			return 0
		}
		if grid[row][col] == 1 {
			return 1
		} else {
			return 0
		}
	}
}
