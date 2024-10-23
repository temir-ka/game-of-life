package main

import (
	. "crunch03/colors"
	_ "crunch03/errors"
	. "crunch03/flags"
	. "crunch03/game"
	_ "crunch03/helpers"
	. "crunch03/input"
	. "crunch03/messages"
	"fmt"
)

func main() {
	err := ""
	options := map[string]string{}
	options, err = GetFlags()

	grid := [][]int{}
	var height, width int
	if err != "" {
		ErrorMessage(Color(err, "", "red"))
		return
	}
	if HelpMode(options) {
		return
	}

	if height, width, grid, err = FileReadingMode(options); err == "SUCCESS" {
	
		// Successfully read from file
	} else if len(RandomMode(options)) > 0 {
		size := RandomMode(options)
		height, width = size[0], size[1]
		grid = RandomMap(height, width)
	} else if len(FullscreenMode(options)) > 0 {
		size := FullscreenMode(options)
		height, width = size[0], size[1]/2
		grid = RandomMap(height, width)

	} else {
		fmt.Println(err)
		height, width, grid, err = Input()
		if err != "" {
			ErrorMessage(err)
			return
		}
	}

	Run(height, width, grid, options)
}
