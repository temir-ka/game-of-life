package flags

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func GetFlags() (map[string]string, string) {
	allFlags := map[string]bool{
		"--help":    true,
		"--verbose": true,
		//"--delay-ms=":  true,
		//"--random=":    true,
		"--footprints":   true,
		"--edges-portal": true,
		"--colored":      true,
		"--fullscreen":   true,
	}
	options := os.Args[1:]
	flags, err := RecordError(allFlags, options)
	if err != "" {
		return map[string]string{}, err
	}
	err = ConflictError(flags)
	if err != "" {
		return map[string]string{}, err
	}
	return flags, ""
}

func RecordError(allFlags map[string]bool, options []string) (map[string]string, string) {
	flags := map[string]string{}
	for _, val := range options {
		if !allFlags[val] {
			stop := true
			var key, str string
			key, str = CheckFlagDelay(val)
			if str == "" && key != "" {
				return map[string]string{}, key
			} else if key == "--delay-ms=" {
				stop = false
				_, ok := flags[key]
				if !ok {
					flags[key] = str
				}
			}
			key, str = CheckFlagRandom(val)
			if str == "" && key != "" {
				return map[string]string{}, key
			} else if key == "--random=" {
				stop = false
				_, ok := flags[key]
				if !ok {
					flags[key] = str
				}
			}
			key, str = CheckFlagFile(val)
			if str == "" && key != "" {
				return map[string]string{}, key
			} else if key == "--file=" {
				stop = false
				_, ok := flags[key]
				if !ok {
					flags[key] = str
				}
			}
			if stop {
				errorMessage := fmt.Sprintf("Option '%s' not found", val)
				return map[string]string{}, errorMessage
			}

		} else {
			_, ok := flags[val]
			if !ok {
				flags[val] = ""
			}
		}
	}
	return flags, ""
}

func ConflictError(flags map[string]string) string {
	_, ok := flags["--help"]
	if ok {
		if len(flags) > 1 {
			return "Option '--help' conflict with other options"
		} else {
			return ""
		}
	}
	_, ok2 := flags["--fullscreen"]
	if ok2 {
		if len(flags) == 1 {
			return ""
		} else {
			for eFlag, _ := range flags {
				if eFlag == "--random=" || eFlag == "--file=" || eFlag == "--verbose" {
					return "Option '--fullscreen' conflict with " + eFlag + " option"
				}
			}
		}
	}
	_, ok3 := flags["--random="]
	if ok3 {
		if len(flags) == 1 {
			return ""
		} else {
			for eFlag, _ := range flags {
				if eFlag == "--file=" {
					return "Option '--random' conflict with " + eFlag + " option"
				}
			}
		}
	}

	return ""
}

func CheckFlagDelay(s string) (string, string) {
	pattern := "^--delay-ms=.*$"
	re := regexp.MustCompile(pattern)
	errorMessage := ""
	if re.MatchString(s) {
		value := s[11:] // check for num
		delay, err := strconv.Atoi(value)
		if err != nil {
			errorMessage = "Value of '--delay-ms=' must be natural number"
			return errorMessage, ""
		}
		if delay < 1 || delay > 10000 {
			errorMessage = "Value of '--delay-ms=' must be in range (1, 10000)"
			return errorMessage, ""
		}
		return "--delay-ms=", strconv.Itoa(delay)
	}
	return "", ""
}

func CheckFlagRandom(s string) (string, string) {
	pattern := "^--random=.*$"
	re := regexp.MustCompile(pattern)
	errorMessage := ""
	if re.MatchString(s) {
		value := s[9:] // check for num
		x_index := -1
		for i := 0; i < len(value); i++ {
			if value[i] == 'x' {
				x_index = i
				break
			}
		}
		if x_index == -1 {
			errorMessage = "Value of '--random' must be in format HxW"
			return errorMessage, ""
		}
		h := value[:x_index]
		w := value[x_index+1:]
		height, errH := strconv.Atoi(h)
		if errH != nil {
			errorMessage = "Value of height of '--random' must be natural number"
			return errorMessage, ""
		}
		width, errW := strconv.Atoi(w)
		if errW != nil {
			errorMessage = "Value of width of '--random' must be natural number"
			return errorMessage, ""
		}
		if height < 3 || height > 73 || width < 3 || width > 185 {
			errorMessage = "Values of height and width must be in range x:[3, 185] y: [3, 73]"
			return errorMessage, ""
		}
		return "--random=", value
	}
	return "", ""
}

func CheckFlagFile(s string) (string, string) {
	pattern := "^--file=.*$"
	re := regexp.MustCompile(pattern)
	errorMessage := ""
	if re.MatchString(s) {
		filePath := s[7:] // Extract the file path
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			errorMessage = fmt.Sprintf("File '%s' does not exist", filePath)
			return errorMessage, ""
		}
		return "--file=", filePath
	}
	return "", ""
}
