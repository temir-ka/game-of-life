package game

import (
	. "crunch03/messages"
	"strconv"
)

func SetDelay(options map[string]string) int {
	delay := 2500
	delay_string, ok := options["--delay-ms="]
	if ok {
		delay, _ = strconv.Atoi(delay_string)
	}
	return delay
}

func TurnModes(options map[string]string) map[string]bool {
	modes := map[string]bool{
		"--verbose":      false,
		"--footprints":   false,
		"--edges-portal": false,
		"--colored":      false,
		"--fullscreen":   false,
	}
	for key, _ := range options {
		_, ok := modes[key]
		if ok {
			modes[key] = true
		}
	}
	return modes
}

func VerboseMode(turned bool, tick, height, width, aliveCells, delay int) {
	if !turned {
		return
	}
	VerboseMessage(tick, height, width, aliveCells, delay)
}
