package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	h      string
	w      string
	input  string
	height int
	width  int
)

func Input() (int, int, [][]int, string) {
	fmt.Printf("Enter the height and width of the map: ")
	fail := make([][]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	var arrF int
	if scanner.Scan() {
		line := scanner.Text()
		arr := strings.Fields(line)
		arrF = len(arr)
		if len(arr) == 2 {
			h = arr[0]
			w = arr[1]
			if _, err := strconv.Atoi(h); err != nil {
				return 0, 0, fail, h + " is not a number"
			}
			if _, err := strconv.Atoi(w); err != nil {
				return 0, 0, fail, w + " is not a number"
			}
			res := fail
			height, _ := strconv.Atoi(h)
			width, _ := strconv.Atoi(w)
			if height < 3 || height > 73 || width < 3 || width > 185 {
				return 0, 0, fail, h + " x " + w + ": incorrect size\nInput range is [3, 73] x [3, 185]"
			}
			dead := 0
			live := 0
			for i := 0; i < height; i++ {
				values := make([]int, 0)
				fmt.Scanf("%s", &input)
				if len(input) < width {
					return 0, 0, fail, "You entered only " + strconv.Itoa(len(input)) + " elements from " + w
				}
				if len(input) > width {
					return 0, 0, fail, "You entered " + strconv.Itoa(len(input)) + " elements from " + w
				}
				for _, eRune := range input {
					if eRune != '.' && eRune != '#' {
						return 0, 0, fail, "Illegal symbol: " + string(eRune)
					}
					if eRune == '.' {
						dead++
						values = append(values, 0)
					}
					if eRune == '#' {
						live++
						values = append(values, 1)
					}
				}

				res = append(res, values)

			}
			if live < 3 {
				return 0, 0, fail, "Write at least 3 live cells"
			}
			if dead < 3 {
				return 0, 0, fail, "Dead cells can't be less than 3"
			}
			return height, width, res, ""
		}
	}
	return 0, 0, fail, "You entered " + strconv.Itoa((arrF)) + " values out of 2 (height and width)"
}
