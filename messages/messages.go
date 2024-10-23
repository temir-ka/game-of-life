package messages

import (
	. "crunch03/colors"
	"fmt"
)

func HelpMessage() {
	fmt.Printf("Usage: go run main.go [options]\n\n")
	fmt.Println("Options:")
	fmt.Printf("--help\t\t: Show the help message and exit\n")
	fmt.Printf("--verbose\t: Display detailed information about the simulation, including grid size, number of ticks, speed, and map name\n")
	fmt.Printf("--delay-ms=X\t: Set the animation speed in milliseconds. Default is 2500 milliseconds\n")
	fmt.Printf("--file=X\t: Load the initial grid from a specified file\n")
	fmt.Printf("--edges-portal\t: Enable portal edges where cells that exit the grid appear on the opposite side\n")
	fmt.Printf("--random=WxH\t: Generate a random grid of the specified width (W) and height (H)\n")
	fmt.Printf("--fullscreen\t: Adjust the grid to fit the terminal size with empty cells\n")
	fmt.Printf("--footprints\t: Add traces of visited cells, displayed as '∘'\n")
	fmt.Printf("--colored\t: Add color to live cells and traces if footprints are enabled\n")
}

func VerboseMessage(tick, height, width, aliveCells, delay int) {
	fmt.Printf("Tick: %d\n", tick)
	fmt.Printf("Grid Size: %dx%d\n", height, width)
	fmt.Printf("Live Cells: %d\n", aliveCells)
	fmt.Printf("DelayMs: %dms\n\n", delay)
}

func ErrorMessage(err string) {
	fmt.Println(Color(err, "", "red"))
}

func StopMessage() {
	fmt.Println("Game Over: Reproduction is no longer possible.")
	fmt.Println("Thank you for playing!")
}
