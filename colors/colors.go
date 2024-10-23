package colors

func Color(text, fontColor, textColor string) string {
	var (
		font         string
		color        string
		defaultColor string = "\033[0m"
	)
	switch fontColor {
	case "red":
		font = "\033[41m"
	case "green":
		font = "\033[42m"
	case "blue":
		font = "\033[44m"
	case "ocean":
		font = "\033[46m"
	case "yellow":
		font = "\033[43m"
	case "violet":
		font = "\035[43m"
	default:
		font = "\033[0m"

	}

	switch textColor {
	case "red":
		color = "\033[31m"
	case "green":
		color = "\033[32m"
	case "blue":
		color = "\033[34m"
	case "ocean":
		color = "\033[36m"
	case "yellow":
		color = "\033[33m"
	case "violet":
		font = "\035[33m"
	default:
		color = "\033[30m"

	}
	return font + color + text + defaultColor
}
