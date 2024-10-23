package flags

import (
	. "crunch03/file_mode"
	. "crunch03/messages"
	"strconv"

	tsize "github.com/kopoli/go-terminal-size"
)

func HelpMode(options map[string]string) bool {
	_, help := options["--help"]
	if help {
		HelpMessage()
		return true
	}
	return false
}

func RandomMode(options map[string]string) []int {
	val, random := options["--random="]
	if random {
		x_index := 0
		for i := 0; i < len(val); i++ {
			if val[i] == 'x' {
				x_index = i
			}
		}
		size := make([]int, 2)
		size[0], _ = strconv.Atoi(val[:x_index])
		size[1], _ = strconv.Atoi(val[x_index+1:])
		return size
	}
	return []int{}
}

func FullscreenMode(options map[string]string) []int {
	_, fullscreen := options["--fullscreen"]
	if fullscreen {
		var s tsize.Size

		s, err2 := tsize.GetSize()
		if err2 == nil {
			size := make([]int, 2)
			xx := s.Width
			yy := s.Height
			if xx < 3 {
				xx = 3
			}
			if yy < 4 {
				yy = 4
			}
			size[0] = yy - 1
			size[1] = xx
			return size
		}

	}
	return []int{}
}

func FileReadingMode(options map[string]string) (int, int, [][]int, string) {
	if fileFlag, exists := options["--file="]; exists {
		height, width, grid, err := ReadGridFromFile(fileFlag)
		if err != "SUCCESS" {
			return 0, 0, nil, err
		}
		return height, width, grid, "SUCCESS"
	}
	return 0, 0, nil, ""
}
