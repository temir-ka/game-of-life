package flags

import (
	"bufio"
	"fmt"
	"os"
	//"regexp"
	//"strconv"
)

func ReadGridFromFile(filePath string) (int, int, [][]int, string) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return 0, 0, nil, fmt.Sprintf("Could not open file: %s", filePath)
	}
	defer file.Close()

	var grid [][]int
	var height, width int

	
	scanner := bufio.NewScanner(file)

	
	for scanner.Scan() {
		line := scanner.Text()

		
		if width == 0 {
			
			width = len(line)
		} else if len(line) != width {
			return 0, 0, nil, "Inconsistent row lengths in the grid from file"
		}

		row := make([]int, len(line))
		for i, char := range line {
			switch char {
			case '.':
				row[i] = 0
			case '#':
				row[i] = 1 
			default:
				return 0, 0, nil, fmt.Sprintf("Illegal symbol '%c' in grid from file", char)
			}
		}
		grid = append(grid, row)
	}


	if err := scanner.Err(); err != nil {
		return 0, 0, nil, fmt.Sprintf("Error reading file: %s", err)
	}

	height = len(grid)

	
	if height < 3 || width < 3 {
		return 0, 0, nil, fmt.Sprintf("Grid size must be at least 3x3, but got %dx%d", height, width)
	}

	return height, width, grid, "SUCCESS"
}
