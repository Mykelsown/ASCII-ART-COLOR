package asciiart

import (
	"strings"
	"log"
)

var colorStorage = map[string]string{
	"red":     "\033[38;2;255;0;0m",
	"green":   "\033[38;2;0;255;0m",
	"yellow":  "\033[38;2;255;255;0m",
	"blue":    "\033[38;2;0;0;255m",
	"magenta": "\033[38;2;255;0;255m",
	"cyan":    "\033[38;2;0;255;255m",
	"white":   "\033[38;2;255;255;255m",
	"black":   "\033[38;2;0;0;0m",

	"bright_red":     "\033[38;2;255;85;85m",
	"bright_green":   "\033[38;2;85;255;85m",
	"bright_yellow":  "\033[38;2;255;255;85m",
	"bright_blue":    "\033[38;2;85;85;255m",
	"bright_magenta": "\033[38;2;255;85;255m",
	"bright_cyan":    "\033[38;2;85;255;255m",
	"bright_white":   "\033[38;2;255;255;255m", // Usually same as white
}


func ApplyColor(colorType string, arguments []string) string {

	//Checks if the color user passes in is not in the list of color available
	if colorStorage[colorType] == "" {
		log.Fatalf("the color (%v) you passed isn't in the list of colors available", colorType)
	}

	mainString := arguments[len(arguments)-1]
	options := arguments[:len(arguments)-1]

	formattedMain := strings.Split(FormatPrinter(mainString), "\n")

	color := colorStorage[colorType]
	reset := "\033[0m"

	// If no substring → color everything
	if len(options) == 0 {
		var result strings.Builder
		for _, line := range formattedMain {
			result.WriteString(color + line + reset + "\n")
		}
		return result.String()
	}

	// Prepare result lines
	resultLines := make([]string, len(formattedMain))

	// Walk through string (instead of Cut)
	for i := 0; i < len(mainString); {
		matched := false

		for _, sub := range options {
			if sub != "" && strings.HasPrefix(mainString[i:], sub) {
				asciiSub := strings.Split(FormatPrinter(sub), "\n")

				for line := range resultLines {
					resultLines[line] += color + asciiSub[line] + reset
				}

				i += len(sub)
				matched = true
				break
			}
		}

		if !matched {
			char := string(mainString[i])
			asciiChar := strings.Split(FormatPrinter(char), "\n")

			for line := range resultLines {
				resultLines[line] += asciiChar[line]
			}
			i++
		}
	}

	return strings.Join(resultLines, "\n")
}