package asciiart

import (
	"fmt"
	"strings"
)

var colorStorage = map[string]string{
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
	"black":   "\033[30m",

	// Bright variants
	"bright_red":     "\033[91m",
	"bright_green":   "\033[92m",
	"bright_yellow":  "\033[93m",
	"bright_blue":    "\033[94m",
	"bright_magenta": "\033[95m",
	"bright_cyan":    "\033[96m",
	"bright_white":   "\033[97m",
}

func ApplyColor(colorType string, arguments []string) string {
	mainString := arguments[len(arguments)-1] // gets the main string in the argument
	options := arguments[:len(arguments)-1]   // gets the substrings(s) that needs to be colored in the main string from the argument

	formatedMainString := FormatPrinter(mainString) //This gets the ascii-art of the main string
	// This is the main string that will store the final output
	var newFormattedMain strings.Builder

	// The main logic that will check for each sub-string in the main string and then apply the color to them
	for _, v := range options {
		formatedOption := FormatPrinter(v) // This stores the sub-string ascii-art
		// Checks if the substring passed from the terminal really exist in the main string
		contains := strings.Contains(mainString, v)

		if contains {
			coloredFormattedOption := colorStorage[colorType] + formatedOption + "\033[0m"

			splittedFormattedMainString := strings.Split(formatedMainString, "\n")
			splittedFormattedOption := strings.Split(formatedOption, "\n")
			splittedColoredFormattedOption := strings.Split(coloredFormattedOption, "\n")

			var slicedModifiedMainString = make([]string, len(splittedFormattedMainString))

			for i, asciiChars := range splittedFormattedMainString {
				slicedModifiedMainString[i] = strings.Replace(asciiChars, splittedFormattedOption[i], splittedColoredFormattedOption[i], 1)
				// fmt.Println(slicedModifiedMainString)
			}
			fmt.Println(strings.Join(slicedModifiedMainString, "\n"))
			newFormattedMain.WriteString(formatedMainString)
		}
	}

	return newFormattedMain.String()
}
